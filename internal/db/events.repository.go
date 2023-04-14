package db

import (
	"context"
	"log"

	"github.com/redundant4u/DeoDeokGo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventsRepository interface {
	Add(ctx context.Context, e models.Event) ([]byte, error)
	Find(ctx context.Context, id string) (models.Event, error)
	FindByName(ctx context.Context, name string) (models.Event, error)
	FindAll(ctx context.Context) ([]models.Event, error)
}

type EventsCollection struct {
	collection *mongo.Collection
}

func NewEventsRepository(db MongoDatabase) EventsRepository {
	return &EventsCollection{collection: db.Collection("events")}
}

func (c *EventsCollection) Add(ctx context.Context, e models.Event) ([]byte, error) {
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}

	_, err := c.collection.InsertOne(ctx, e)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(e.ID.Hex()), nil
}

func (c *EventsCollection) Find(ctx context.Context, id string) (models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Event{}, err
	}

	filter := bson.M{"_id": objectId}

	result := c.collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return models.Event{}, err
	}

	var e models.Event
	if err := result.Decode(&e); err != nil {
		return models.Event{}, nil
	}

	return e, nil
}

func (c *EventsCollection) FindByName(ctx context.Context, name string) (models.Event, error) {
	filter := bson.M{"name": name}

	result := c.collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return models.Event{}, err
	}

	var e models.Event
	if err := result.Decode(&e); err != nil {
		return models.Event{}, err
	}

	return e, nil
}

func (c *EventsCollection) FindAll(ctx context.Context) ([]models.Event, error) {
	filter := bson.M{}
	options := options.Find()

	result, err := c.collection.Find(ctx, filter, options)

	if err != nil {
		return []models.Event{}, err
	}

	var e []models.Event
	if err := result.All(ctx, &e); err != nil {
		return []models.Event{}, err
	}

	return e, nil
}
