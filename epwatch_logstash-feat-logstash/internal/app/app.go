package app

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/engine"
	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/repository"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	log.Println("LOGSTASH_OUTPUT:", os.Getenv("LOGSTASH_OUTPUT"))
	log.Println("CONTAINER_ID:", os.Getenv("CONTAINER_ID"))
	// path := viper.GetString("file.path")
	mongoAddress := viper.GetString("mongodb.address")
	mongoDb := viper.GetString("mongodb.database")
	logSources := viper.GetStringSlice("file.sources")
	blobAddress := viper.GetString("blob.connection_string")
	mapCollectionName := viper.GetString("mongodb.map_collections")

	logStashOutput := os.Getenv("LOGSTASH_OUTPUT")
	containerId := os.Getenv("CONTAINER_ID")
	path := filepath.Join(logStashOutput, containerId)

	fileRepo := repository.NewLogstashFileRepository(path)
	err := fileRepo.CreateDirIfNotExists()
	if err != nil {
		log.Println(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoAddress))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	serviceClient, err := azblob.NewServiceClientFromConnectionString(blobAddress, nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	mapCollection := client.Database(mongoDb).Collection(mapCollectionName)
	var wg sync.WaitGroup
	for _, source := range logSources {
		wg.Add(1)
		fileRepo := repository.NewLogstashFileRepository(filepath.Join(path, source))
		collection := client.Database(mongoDb).Collection(source)
		mongoRepo := repository.NewLogStashMongoRepository(client, collection, mongoDb, source, mapCollection)
		// publicAccessType := azblob.PublicAccessTypeContainer
		containerCreateOptions := azblob.ContainerCreateOptions{}
		_, err := serviceClient.CreateContainer(ctx, source, &containerCreateOptions)
		if err != nil {
			log.Println(err)
		}
		containerClient, err := serviceClient.NewContainerClient(source)
		if err != nil {
			log.Println(err)
		}
		blobRepository := repository.NewLogstashBlobRepository(containerClient)
		engine := engine.NewLogstashEngine(mongoRepo, blobRepository, fileRepo)
		ctx = context.TODO()
		go engine.ScanLogstash(ctx, &wg)
	}
	wg.Wait()
}
