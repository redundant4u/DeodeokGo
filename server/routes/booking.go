package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/controllers/booking"
	"github.com/redundant4u/DeoDeokGo/controllers/events"
	"github.com/redundant4u/DeoDeokGo/db"
	"github.com/redundant4u/DeoDeokGo/queue"
)

func InitBookingRoutes(ctx context.Context, db db.MongoDatabase, listener queue.EventListener, emitter queue.EventEmitter) *gin.Engine {
	r := gin.Default()

	eventsRepository := events.NewRepository(ctx, db)

	bookingRepository := booking.NewRepository(ctx, db)
	bookingService := booking.NewService(bookingRepository, eventsRepository, emitter)
	bookingController := booking.NewController(bookingService)

	processor := &booking.Processor{
		EventsRepository: eventsRepository,
		Listener:         listener,
	}
	go processor.ProcessEvents()

	bookingGroup := r.Group("booking")

	bookingGroup.POST("/events/:eventID", bookingController.NewBooking)

	return r
}
