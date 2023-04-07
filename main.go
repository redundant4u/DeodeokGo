package main

import (
	"context"

	"github.com/redundant4u/DeoDeokGo/internal/config"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/server"
)

func main() {
	ctx := context.Background()
	env := "dev"
	cfg := config.LoadConfig(env)

	client := db.InitMongoDBClient(ctx, cfg)

	server.Init(client)
}