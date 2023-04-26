package main

import (
	"context"
	"log"

	"github.com/redundant4u/DeoDeokGo/config"
	"github.com/redundant4u/DeoDeokGo/db"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/amqp"
	"github.com/redundant4u/DeoDeokGo/queue/kafka"
	"github.com/redundant4u/DeoDeokGo/routes"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatal("Colud not load config: ", err)
	}

	// DB
	mongoClient, err := db.NewMongoClient(cfg)
	if err != nil {
		log.Fatal("Colud not connect DB: ", err)
	}

	// Message Broker
	amqpConn, err := amqp.NewAmqpConnection(cfg)
	if err != nil {
		log.Fatal("Could not establish AMQP connection: ", err)
	}
	defer amqpConn.Close()

	var eventListener queue.EventListener
	var eventEmitter queue.EventEmitter

	switch cfg.GetString("broker.type") {
	case "amqp":
		eventEmitter = amqp.NewEventEmitter(ctx, amqpConn, "events")
		eventListener = amqp.NewEventListener(amqpConn, "events")
	case "kafka":
		kafkaClient := kafka.NewKafkaClient(cfg)
		eventEmitter = kafka.NewKafkaEventEmitter(kafkaClient)
		eventListener = kafka.NewKafkaEventListener(kafkaClient, []int32{})
	default:
		panic("Bad broker type name")
	}

	// Router
	r := routes.InitEventsRoutes(ctx, mongoClient.Database(), eventListener, eventEmitter)

	log.Fatal(r.Run())
}
