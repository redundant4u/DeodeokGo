package db

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDBClient(ctx context.Context, cfg *viper.Viper) *mongo.Client {
	url := cfg.GetString("database.uri")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(fmt.Errorf("Could not connect to database: %w", err))
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(fmt.Errorf("Cloud not ping database: %w", err))
	}

	return client
}
