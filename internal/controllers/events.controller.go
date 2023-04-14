package controllers

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/db"
	"github.com/redundant4u/DeoDeokGo/internal/models"
	"github.com/redundant4u/DeoDeokGo/internal/queue"
	"github.com/redundant4u/DeoDeokGo/internal/queue/contracts"
)

type eventsController struct {
	repository db.EventsRepository
	emitter    queue.EventEmitter
}

func NewEventsController(repository db.EventsRepository, emitter queue.EventEmitter) *eventsController {
	c := &eventsController{
		repository: repository,
		emitter:    emitter,
	}

	return c
}

func (c *eventsController) NewEvent(ctx *gin.Context) {
	event := models.Event{}

	if err := ctx.BindJSON(&event); err != nil {
		log.Fatal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	fmt.Println(event)

	id, err := c.repository.Add(ctx, event)

	if err != nil {
		log.Fatal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: event.Location.ID.String(),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}

	err = c.emitter.Emit(&msg)
	fmt.Println(err)

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *eventsController) FindAllEvents(ctx *gin.Context) {
	events, err := c.repository.FindAll(ctx)

	if err != nil {
		log.Default().Fatal("FindAllEvents Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, events)
}

func (c *eventsController) FindEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := c.repository.Find(ctx, id)

	if err != nil {
		log.Default().Fatal("FindEvent Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *eventsController) FindEventByName(ctx *gin.Context) {
	name := ctx.Param("name")

	event, err := c.repository.FindByName(ctx, name)

	if err != nil {
		log.Default().Fatal("FindEventByName Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}
