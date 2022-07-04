package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/domain"
	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/repository"
)

type LogstashEngine struct {
	logstashMongoRepo *repository.LogStashMongoRepository
	logstashBlobRepo  *repository.LogstashBlobRepository
	logstashFileRepo  *repository.LogstashFileRepository
}

func NewLogstashEngine(
	logstashMongoRepo *repository.LogStashMongoRepository,
	logstashBlobRepo *repository.LogstashBlobRepository,
	logstashFileRepo *repository.LogstashFileRepository,
) *LogstashEngine {
	return &LogstashEngine{
		logstashMongoRepo: logstashMongoRepo,
		logstashBlobRepo:  logstashBlobRepo,
		logstashFileRepo:  logstashFileRepo,
	}
}

func (engine *LogstashEngine) Push(ctx context.Context, logName string, logTextBatch []string) error {
	log.Printf("Push %s lenght %d\n", logName, len(logTextBatch))
	if len(logTextBatch) == 0 {
		return nil
	}
	var mongoBatch []interface{}
	for _, logText := range logTextBatch {
		var jsonMap map[string]interface{}
		err := json.Unmarshal([]byte(logText), &jsonMap)
		if err != nil {
			log.Println(err)
		}
		_, ok1 := jsonMap["dst_hostname"]
		_, ok2 := jsonMap["username"]
		if ok1 || ok2 {
			mongoBatch = append(mongoBatch, jsonMap)
		}
	}
	if len(mongoBatch) > 0 {
		log.Printf("Upload mognodb %s lenght %d\n", logName, len(mongoBatch))
		err := engine.logstashMongoRepo.InsertMany(ctx, mongoBatch)
		if err != nil {
			log.Println("InsertMany", err)
			return err
		}
	}
	// send to blob
	timestamp := time.Now().Format(time.RFC3339Nano)
	logName = fmt.Sprintf("%s-%s", timestamp, logName)
	log.Printf("Upload blob %s lenght %d\n", logName, len(logTextBatch))
	blobClient, err := engine.logstashBlobRepo.UploadBlob(ctx, logName, logTextBatch)
	if err != nil {
		log.Println(err)
		return err
	}
	f := "2006-01-02"
	t := time.Now().Format(f)
	TimeStamp, _ := time.Parse(f, t)
	blobMapper := domain.BlobMapper{
		BlobUrl:   blobClient.URL(),
		TimeStamp: TimeStamp,
	}
	err = engine.logstashMongoRepo.InserBlobMapperOne(ctx, blobMapper)
	if err != nil {
		log.Println("InsertOne", err)
		return err
	}
	return nil
}

func (engine *LogstashEngine) Worker(ctx context.Context, logName string) error {
	logChan := make(chan string)
	quitChan := make(chan bool)
	lineNumberChan := make(chan int)
	defer close(logChan)
	defer close(quitChan)
	defer close(lineNumberChan)
	go engine.logstashFileRepo.GetLog(logName, 0, logChan, quitChan)
	var wg sync.WaitGroup
	var logTextBatch []string
	var lineNumber int
	var startTime time.Time = time.Now().Add(5 * time.Minute)

	go func() {
		for lineNumber := range lineNumberChan {
			// save index
			log.Printf("Upload check point %s line number %d\n", logName, lineNumber)
			err := engine.logstashFileRepo.SaveCheckPoint(logName, lineNumber)
			if err != nil {
				log.Println("Worker", err)
			}
		}
	}()
L:
	for {
		select {
		case <-quitChan:
			err := engine.Push(ctx, logName, logTextBatch)
			if err != nil {
				log.Println("Worker", err)
			}
			// delete file
			wg.Wait()
			err = engine.logstashFileRepo.DeleteFile(logName)
			if err != nil {
				log.Println("Worker", err)
			}
			break L
		case logText := <-logChan:
			logTextBatch = append(logTextBatch, logText)
			if len(logTextBatch) >= 100000 || startTime.Before(time.Now()) {
				// send to mongodb
				wg.Add(1)
				go func(logTextBatch []string, lineNumber int) {
					err := engine.Push(ctx, logName, logTextBatch)
					if err != nil {
						log.Println("Worker", err)
					}
					lineNumberChan <- lineNumber
					wg.Done()
				}(logTextBatch, lineNumber)
				// clear batch
				logTextBatch = nil
				startTime = time.Now().Add(5 * time.Minute)
			}
			lineNumber++
		}
	}
	return nil
}

func (engine *LogstashEngine) ScanLogstash(ctx context.Context, wg *sync.WaitGroup) {
	// looking for log file
	err := engine.logstashFileRepo.CreateDirIfNotExists()
	if err != nil {
		log.Println(err)
		return
	}
	defer wg.Done()
	for {
		fileNames, err := engine.logstashFileRepo.ListDir()
		if err != nil {
			log.Println(err)
		}
		var scanLogstashWG sync.WaitGroup
		for _, fileName := range fileNames {
			var extension = filepath.Ext(fileName.Name())
			if extension == ".log" {
				logName := fileName.Name()
				// scanLogstashWG.Add(1)
				// go func() {
				// 	defer scanLogstashWG.Done()
				// 	err := engine.Worker(ctx, logName)
				// 	if err != nil {
				// 		log.Println("ScanLogstash", err)
				// 	}
				// 	log.Println("Worker Done")
				// }()
				err := engine.Worker(ctx, logName)
				if err != nil {
					log.Println("ScanLogstash", err)
				}
			}
		}
		scanLogstashWG.Wait()
		log.Println("Done", engine.logstashFileRepo.FilePath)
		// break
		time.Sleep(time.Minute * 5)
	}
}
