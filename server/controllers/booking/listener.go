package booking

import (
	"log"

	"github.com/redundant4u/DeoDeokGo/controllers/events"
	"github.com/redundant4u/DeoDeokGo/models"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Processor struct {
	EventsRepository events.Repository
	Listener         queue.EventListener
}

func (p *Processor) ProcessEvents() {
	received, errors, err := p.Listener.Listen("eventCreated")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-received:
			p.handleEvent(event)
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
		_, err := p.EventsRepository.Add(models.Event{ID: objectID, Name: e.Name})
		if err != nil {
			log.Fatal("Add Event Error")
		}
	}
}
