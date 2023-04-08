package mocks

import (
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockMongoClient struct{}

type MockMongoDataBase struct{}

func (c *MockMongoClient) Database() db.MongoDatabase {
	return &MockMongoDataBase{}
}

func (c *MockMongoClient) Ping() error {
	return nil
}

func (c *MockMongoClient) Disconnect() error {
	return nil
}

func (db *MockMongoDataBase) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return nil
}
