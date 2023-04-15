package db

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type MongoClient interface {
	Database() MongoDatabase
}

type mongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoClient(cfg *viper.Viper) (MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.GetString("database.uri"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &mongoClient{
		client:   client,
		database: client.Database(cfg.GetString("database.name")),
	}, nil
}

func (m *mongoClient) Database() MongoDatabase {
	return m.database
}
