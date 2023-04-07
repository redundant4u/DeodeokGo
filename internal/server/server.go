package server

import (
	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/config"
	"github.com/redundant4u/DeoDeokGo/internal/controllers"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) {
	config := config.GetConfig()
	router := gin.Default()

	databaseName := config.GetString("database.name")

	eventsRepository := db.NewEventsRepository(client.Database(databaseName))
	eventsController := controllers.NewEventsController(eventsRepository)

	eventsGroup := router.Group("events")

	eventsGroup.GET("", eventsController.FindAllEvents)
	eventsGroup.GET("/id/:id", eventsController.FindEvent)
	eventsGroup.GET("/name/:name", eventsController.FindEventByName)
	eventsGroup.POST("", eventsController.NewEvent)

	router.Run(":8888")
}
