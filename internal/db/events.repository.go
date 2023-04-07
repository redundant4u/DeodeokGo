package db

import (
	"context"
	"fmt"

	"github.com/redundant4u/DeoDeokGo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventsRepository interface {
	AddEvent(ctx context.Context, e models.Event) error
	FindEvent(ctx context.Context, id string) (models.Event, error)
	FindEventByName(ctx context.Context, name string) (models.Event, error)
	FindAllEvents(ctx context.Context) ([]models.Event, error)
}

type Collection struct {
	collection *mongo.Collection
}

func NewEventsRepository(db *mongo.Database) EventsRepository {
	return &Collection{collection: db.Collection("events")}
}

func (c *Collection) AddEvent(ctx context.Context, e models.Event) error {
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}

	_, err := c.collection.InsertOne(ctx, e)

	if err != nil {
		return err
	}

	return nil
}

func (c *Collection) FindEvent(ctx context.Context, id string) (models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	fmt.Println(id, objectId)

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

func (c *Collection) FindEventByName(ctx context.Context, name string) (models.Event, error) {
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

func (c *Collection) FindAllEvents(ctx context.Context) ([]models.Event, error) {
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
