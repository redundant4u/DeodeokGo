package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/models"
)

type EventsController struct {
	repository db.EventsRepository
}

func NewEventsController(repository db.EventsRepository) *EventsController {
	c := &EventsController{
		repository: repository,
	}

	return c
}

func (c *EventsController) NewEvent(ctx *gin.Context) {
	eventRequest := models.Event{}

	if err := ctx.BindJSON(&eventRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err := c.repository.AddEvent(ctx, eventRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *EventsController) FindAllEvents(ctx *gin.Context) {
	events, err := c.repository.FindAllEvents(ctx)

	if err != nil {
		log.Default().Fatal("FindAllEvents Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, events)
}

func (c *EventsController) FindEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := c.repository.FindEvent(ctx, id)

	if err != nil {
		log.Default().Fatal("FindEvent Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *EventsController) FindEventByName(ctx *gin.Context) {
	name := ctx.Param("name")

	event, err := c.repository.FindEventByName(ctx, name)

	if err != nil {
		log.Default().Fatal("FindEventByName Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}
