package booking

import (
	"time"

	"github.com/redundant4u/DeoDeokGo/controllers/booking/dto"
	"github.com/redundant4u/DeoDeokGo/controllers/events"
	"github.com/redundant4u/DeoDeokGo/models"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
)

type Service interface {
	Add(eventID string, b dto.CreateBookingRequest) error
}

type service struct {
	bookingRepository Repository
	eventsRepository  events.Repository
	emitter           queue.EventEmitter
}

func NewService(bookingRepository Repository, eventsRepository events.Repository, emitter queue.EventEmitter) *service {
	return &service{
		bookingRepository: bookingRepository,
		eventsRepository:  eventsRepository,
		emitter:           emitter,
	}
}

func (s *service) Add(eventID string, b dto.CreateBookingRequest) error {
	event, err := s.eventsRepository.Find(eventID)
	if err != nil {
		return err
	}

	eventIDAsBytes, _ := event.ID.MarshalText()
	booking := models.Booking{
		Date:    time.Now().Unix(),
		EventID: eventIDAsBytes,
		Seats:   b.Seats,
	}

	err = s.bookingRepository.Add(booking)
	if err != nil {
		return err
	}

	msg := contracts.EventBookedEvent{
		EventID:  event.ID.Hex(),
		MemberID: "test",
	}

	err = s.emitter.Emit(&msg)
	if err != nil {
		return err
	}

	return nil
}
