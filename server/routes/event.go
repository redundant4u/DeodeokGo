package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/controllers/events"
	"github.com/redundant4u/DeoDeokGo/db"
	"github.com/redundant4u/DeoDeokGo/queue"
)

func InitEventsRoutes(ctx context.Context, db db.MongoDatabase, emitter queue.EventEmitter) *gin.Engine {
	r := gin.Default()

	eventsRepository := events.NewRepository(ctx, db)
	eventsService := events.NewService(eventsRepository, emitter)
	eventsController := events.NewController(eventsService)

	eventsGroup := r.Group("events")

	eventsGroup.GET("", eventsController.FindAll)
	eventsGroup.GET("/id/:id", eventsController.Find)
	eventsGroup.GET("/name/:name", eventsController.FindByName)
	eventsGroup.POST("", eventsController.NewEvent)

	return r
}
