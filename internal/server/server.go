package server

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/controllers"
	"github.com/redundant4u/DeoDeokGo/internal/db"
)

var runOnce sync.Once

func Init(c db.MongoClient) {
	runOnce.Do(func() {
		router := Router(c)
		router.Run(":8888")
	})
}

func Router(c db.MongoClient) (router *gin.Engine) {
	router = gin.Default()

	database := c.Database()

	eventsRepository := db.NewEventsRepository(database)
	eventsController := controllers.NewEventsController(eventsRepository)

	eventsGroup := router.Group("events")

	eventsGroup.GET("", eventsController.FindAllEvents)
	eventsGroup.GET("/id/:id", eventsController.FindEvent)
	eventsGroup.GET("/name/:name", eventsController.FindEventByName)
	eventsGroup.POST("", eventsController.NewEvent)

	return router
}
