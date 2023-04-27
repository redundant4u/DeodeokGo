package booking

import (
	"context"

	"github.com/redundant4u/DeoDeokGo/db"
	"github.com/redundant4u/DeoDeokGo/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Add(b models.Booking) error
}

type repository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewRepository(ctx context.Context, db db.MongoDatabase) Repository {
	return &repository{
		ctx:        ctx,
		collection: db.Collection("booking"),
	}
}

func (r *repository) Add(b models.Booking) error {
	_, err := r.collection.InsertOne(r.ctx, b)
	if err != nil {
		return err
	}

	return nil
}
