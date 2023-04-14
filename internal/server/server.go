package server

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/controllers"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/queue"
)

var runOnce sync.Once

func Init(eventsRepository db.EventsRepository, emitter queue.EventEmitter) {
	runOnce.Do(func() {
		router := Router(eventsRepository, emitter)
		router.Run(":8888")
	})
}

func Router(eventsRepository db.EventsRepository, emitter queue.EventEmitter) (router *gin.Engine) {
	router = gin.Default()

	eventsController := controllers.NewEventsController(eventsRepository, emitter)

	eventsGroup := router.Group("events")

	eventsGroup.GET("", eventsController.FindAllEvents)
	eventsGroup.GET("/id/:id", eventsController.FindEvent)
	eventsGroup.GET("/name/:name", eventsController.FindEventByName)
	eventsGroup.POST("", eventsController.NewEvent)

	return router
}
