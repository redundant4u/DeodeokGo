package main

import (
	"context"

	"github.com/redundant4u/DeoDeokGo/internal/config"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/listener"
	"github.com/redundant4u/DeoDeokGo/internal/queue"
	"github.com/redundant4u/DeoDeokGo/internal/server"
)

func main() {
	ctx := context.Background()
	env := "dev"
	cfg := config.LoadConfig(env)

	c, _ := db.NewMongoClient(ctx, cfg)

	amqpConn := queue.NewAmqpConnection(cfg)
	defer amqpConn.Close()

	// TODO: oh no
	eventsRepository := db.NewEventsRepository(c.Database())

	eventEmitter := queue.NewEventEmitter(ctx, amqpConn, "events")
	eventListener, err := queue.NewEventListener(amqpConn, "events")
	if err != nil {
		panic(err)
	}

	processor := listener.EventProcessor{
		Context:          ctx,
		EventsRepository: eventsRepository,
		EventListener:    eventListener,
	}
	go processor.ProcessEvents()

	server.Init(eventsRepository, eventEmitter)
}
