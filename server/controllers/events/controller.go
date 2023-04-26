package events

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/models"
)

type eventsController struct {
	service Service
}

func NewController(service Service) *eventsController {
	c := &eventsController{
		service: service,
	}

	return c
}

func (c *eventsController) NewEvent(ctx *gin.Context) {
	event := models.Event{}

	if err := ctx.BindJSON(&event); err != nil {
		log.Fatal("NewEvent Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	_, err := c.service.Add(event)
	if err != nil {
		log.Fatal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *eventsController) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := c.service.Find(id)

	if err != nil {
		log.Fatal("FindEvent Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *eventsController) FindByName(ctx *gin.Context) {
	name := ctx.Param("name")

	event, err := c.service.FindByName(name)

	if err != nil {
		log.Fatal("FindEventByName Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *eventsController) FindAll(ctx *gin.Context) {
	events, err := c.service.FindAll()

	if err != nil {
		log.Fatal("FindAllEvents Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, events)
}
