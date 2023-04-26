package events

import (
	"context"

	"github.com/redundant4u/DeoDeokGo/db"
	"github.com/redundant4u/DeoDeokGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Add(e models.Event) ([]byte, error)
	Find(id string) (models.Event, error)
	FindByName(name string) (models.Event, error)
	FindAll() ([]models.Event, error)
}

type repository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewRepository(ctx context.Context, db db.MongoDatabase) Repository {
	return &repository{
		ctx:        ctx,
		collection: db.Collection("events"),
	}
}

func (r *repository) Add(e models.Event) ([]byte, error) {
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}

	_, err := r.collection.InsertOne(r.ctx, e)

	if err != nil {
		return nil, err
	}

	return []byte(e.ID.Hex()), nil
}

func (r *repository) Find(id string) (models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Event{}, err
	}

	filter := bson.M{"_id": objectId}

	result := r.collection.FindOne(r.ctx, filter)

	if err := result.Err(); err != nil {
		return models.Event{}, err
	}

	var e models.Event

	if err := result.Decode(&e); err != nil {
		return models.Event{}, nil
	}

	return e, nil
}

func (r *repository) FindByName(name string) (models.Event, error) {
	filter := bson.M{"name": name}

	result := r.collection.FindOne(r.ctx, filter)

	if err := result.Err(); err != nil {
		return models.Event{}, err
	}

	var e models.Event
	if err := result.Decode(&e); err != nil {
		return models.Event{}, err
	}

	return e, nil
}

func (r *repository) FindAll() ([]models.Event, error) {
	filter := bson.M{}
	options := options.Find()

	result, err := r.collection.Find(r.ctx, filter, options)

	if err != nil {
		return []models.Event{}, err
	}

	var e []models.Event
	if err := result.All(r.ctx, &e); err != nil {
		return []models.Event{}, err
	}

	return e, nil
}
