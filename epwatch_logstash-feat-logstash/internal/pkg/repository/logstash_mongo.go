package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogStashMongoRepository struct {
	mongoClient   *mongo.Client
	logCollection *mongo.Collection
	mongoDb       string
	source        string
	mapCollection *mongo.Collection
}

func NewLogStashMongoRepository(mongoClient *mongo.Client,
	logCollection *mongo.Collection,
	mongoDb string,
	source string,
	mapCollection *mongo.Collection) *LogStashMongoRepository {
	return &LogStashMongoRepository{
		mongoClient:   mongoClient,
		logCollection: logCollection,
		mongoDb:       mongoDb,
		source:        source,
		mapCollection: mapCollection,
	}
}

func (repo *LogStashMongoRepository) InsertMany(ctx context.Context, data []interface{}) error {
	_, err := repo.logCollection.InsertMany(ctx, data)
	if err != nil {
		log.Fatal("InsertMany", err)
		return err
	}
	return nil
}

func (repo *LogStashMongoRepository) InserBlobMapperOne(ctx context.Context, data interface{}) error {
	_, err := repo.mapCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal("InsertMany", err)
		return err
	}
	return nil
}
