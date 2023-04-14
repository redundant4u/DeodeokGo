package listener

import (
	"context"
	"log"

	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/models"
	"github.com/redundant4u/DeoDeokGo/internal/queue"
	"github.com/redundant4u/DeoDeokGo/internal/queue/contracts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventProcessor struct {
	Context          context.Context
	EventsRepository db.EventsRepository
	EventListener    queue.EventListener
}

func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events")

	received, errors, err := p.EventListener.Listen("eventCreated")
	if err != nil {
		log.Fatal(err)
		return err
	}

	// var forever chan struct{}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("received error while processing msg: %s", err)
		}
	}
	// <-forever

	// return nil
}

func (p *EventProcessor) handleEvent(event queue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created %s", e.ID, e)
		objectID, _ := primitive.ObjectIDFromHex(e.ID)
		p.EventsRepository.Add(p.Context, models.Event{ID: objectID})
	}
}
