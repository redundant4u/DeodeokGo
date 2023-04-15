package events

import (
	"encoding/hex"
	"time"

	"github.com/redundant4u/DeoDeokGo/models"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
)

type Service interface {
	Add(e models.Event) ([]byte, error)
	Find(id string) (models.Event, error)
	FindByName(name string) (models.Event, error)
	FindAll() ([]models.Event, error)
}

type service struct {
	repository Repository
	emitter    queue.EventEmitter
}

func NewService(repository Repository, emitter queue.EventEmitter) *service {
	return &service{
		repository: repository,
		emitter:    emitter,
	}
}

func (s *service) Add(e models.Event) ([]byte, error) {
	event := models.Event{
		Name:      e.Name,
		Duration:  e.Duration,
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
		Location:  e.Location,
	}

	id, err := s.repository.Add(event)
	if err != nil {
		return nil, err
	}

	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: event.Location.ID.String(),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}

	err = s.emitter.Emit(&msg)

	return id, err
}

func (s *service) Find(id string) (models.Event, error) {
	events, err := s.repository.Find(id)
	return events, err
}

func (s *service) FindByName(name string) (models.Event, error) {
	event, err := s.repository.FindByName(name)
	return event, err
}

func (s *service) FindAll() ([]models.Event, error) {
	events, err := s.repository.FindAll()
	return events, err
}
