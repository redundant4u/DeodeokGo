package events

import (
	"log"

	"github.com/redundant4u/DeoDeokGo/models"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Processor struct {
	Service  Service
	Listener queue.EventListener
}

func (p *Processor) ProcessEvents() {
	log.Println("Listening to events")

	received, errors, err := p.Listener.Listen("eventCreated")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("received error while processing msg: %s", err)
		}
	}

}

func (p *Processor) handleEvent(event queue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created %s", e.ID, e)

		objectID, _ := primitive.ObjectIDFromHex(e.ID)
		_, err := p.Service.Add(models.Event{ID: objectID, Name: e.Name})
		if err != nil {
			log.Fatal("Add Event Error")
		}
	}
}
