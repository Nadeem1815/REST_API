package routes

import (
	"net/http"
	"strconv"

	"github.com/Nadeem1815/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not fetch event",
		})
		return
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Could not register user for event.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"massage": "Registerd",
	})
}

func CancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"massage": "could not parse event id.",
	// 	})
	// 	return
	// }

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Could not cancel registration.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"massage": "cancelled !"})

}
