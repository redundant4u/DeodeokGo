package booking

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/controllers/booking/dto"
)

type bookingController struct {
	service Service
}

func NewController(service Service) *bookingController {
	return &bookingController{
		service: service,
	}
}

func (c *bookingController) NewBooking(ctx *gin.Context) {
	booking := dto.CreateBookingRequest{}

	if err := ctx.BindJSON(&booking); err != nil {
		log.Fatal("NewBooking Error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if booking.Seats <= 0 {
		log.Fatal("Seat number must be positive")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Seat number must be positive")
	}

	eventID := ctx.Param("eventID")

	err := c.service.Add(eventID, booking)
	if err != nil {
		log.Fatal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusNoContent, nil)
}
