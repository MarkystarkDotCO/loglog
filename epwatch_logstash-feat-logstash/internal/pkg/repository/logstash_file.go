package repository

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/domain"
	"github.com/nxadm/tail"
)

type LogstashFileRepository struct {
	FilePath string
}

func NewLogstashFileRepository(filePath string) *LogstashFileRepository {
	return &LogstashFileRepository{
		FilePath: filePath,
	}
}

func (repo *LogstashFileRepository) GetLog(logName string, offset int, output chan<- string, quitChan chan<- bool) {
	logPath := filepath.Join(repo.FilePath, logName)
	log.Println(logPath)
	t, err := tail.TailFile(
		logPath, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		panic(err)
	}
	lineCount := 0
	startTime := time.Now().Add(time.Minute * 1)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	quit := make(chan bool)
	defer close(quit)
	go func() {
		for t := range ticker.C {
			if startTime.Before(t) {
				fmt.Println(lineCount)
				files, err := repo.ListDir()
				if err != nil {
					log.Println(err)
				}
				for _, f := range files {
					var extension = filepath.Ext(f.Name())
					// There is atleast one log file newer than this.
					if extension == ".log" && f.Name() > logName {
						quit <- true
						return
					}
				}
			}
		}
	}()
L:
	for {
		select {
		case line := <-t.Lines:
			if lineCount > offset {
				output <- line.Text
			}
			startTime = time.Now().Add(time.Minute * 1)
			lineCount++
		case q := <-quit:
			if q {
				quitChan <- true
				break L
			}
		}
	}
}

func (repo *LogstashFileRepository) DeleteFile(logName string) error {
	log.Println("DeleteFile")
	logPath := filepath.Join(repo.FilePath, logName)
	err := os.Remove(logPath)
	if err != nil {
		log.Println(err)
		// return err
	}
	checkPointName := fmt.Sprintf("checkpoint-%s.json", logName[:len(logName)-len(filepath.Ext(logName))])
	checkPointPath := filepath.Join(repo.FilePath, checkPointName)
	err = os.Remove(checkPointPath)
	if err != nil {
		log.Println(err)
		// return err
	}
	return nil
}

func (repo *LogstashFileRepository) OpenCheckPoint(checkPointName string) (map[string]domain.Checkpoint, error) {
	checkPointPath := filepath.Join(repo.FilePath, checkPointName)
	content, err := ioutil.ReadFile(checkPointPath)
	if err != nil {
		log.Println("OpenCheckPoint", err, content)
		return nil, err
	}
	var checkpoints map[string]domain.Checkpoint
	err = json.Unmarshal(content, &checkpoints)
	if err != nil {
		log.Println("OpenCheckPoint", err, content)
		return nil, err
	}
	return checkpoints, nil
}

func (repo *LogstashFileRepository) SaveCheckPoint(logName string, lineNumber int) error {
	checkPointName := fmt.Sprintf("checkpoint-%s.json", logName[:len(logName)-len(filepath.Ext(logName))])
	checkpoint, err := repo.OpenCheckPoint(checkPointName)
	if err != nil {
		log.Println("SaveCheckPoint", err)
	}
	if checkpoint == nil || err != nil {
		checkpoint = make(map[string]domain.Checkpoint)
	}
	if _, ok := checkpoint[logName]; ok {
		if lineNumber < checkpoint[logName].LineNumber {
			lineNumber = checkpoint[logName].LineNumber
		}
	}
	checkpoint[logName] = domain.Checkpoint{
		FileName:   logName,
		LineNumber: lineNumber,
		LastSeen:   time.Now(),
	}
	checkPointPath := filepath.Join(repo.FilePath, checkPointName)
	content, err := json.Marshal(checkpoint)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(checkPointPath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (repo *LogstashFileRepository) ListDir() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(repo.FilePath)
	if err != nil {
		log.Fatal("ListDir", err)
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})
	return files, nil
}

func (repo *LogstashFileRepository) CreateDirIfNotExists() error {
	if _, err := os.Stat(repo.FilePath); os.IsNotExist(err) {
		err := os.Mkdir(repo.FilePath, os.ModePerm)
		if err != nil {
			log.Println("CreateDirIfNotExists", err)
			return err
		}
	}
	return nil
}
