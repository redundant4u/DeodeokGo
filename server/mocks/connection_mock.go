package mocks

import (
	"github.com/redundant4u/DeoDeokGo/queue"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockMongoDatabase struct{}

type MockEventEmitter struct{}

type MockEventListener struct{}

func (db *MockMongoDatabase) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return nil
}

func (e *MockEventEmitter) Emit(event queue.Event) error {
	return nil
}

func (l *MockEventListener) Listen(eventNames ...string) (<-chan queue.Event, <-chan error, error) {
	return nil, nil, nil
}
