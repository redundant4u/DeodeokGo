package db

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type MongoClient interface {
	Database() MongoDatabase
	Ping() error
	Disconnect() error
}

type mongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoClient(ctx context.Context, cfg *viper.Viper) (MongoClient, error) {
	clientOptions := options.Client().ApplyURI(cfg.GetString("database.uri"))
	client, err := mongo.Connect(ctx, clientOptions)

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

func (m *mongoClient) Ping() error {
	if err := m.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("unable to connect to DB: ", err)
		return err
	}
	return nil
}

func (m *mongoClient) Disconnect() error {
	if err := m.client.Disconnect(context.Background()); err != nil {
		log.Fatal("unable to disconnect from DB: ", err)
		return err
	}

	return nil
}
