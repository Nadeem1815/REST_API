package routes

import (
	"net/http"
	"strconv"

	"github.com/Nadeem1815/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not fetch events. try again later",
		})
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {

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
			"massage": "could not fetch events. try again later",
		})

	}
	ctx.JSON(http.StatusOK, event)

}

func CreateEvent(ctx *gin.Context) {
	// token := ctx.Request.Header.Get("Authorization")

	// if token == "" {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"massage": "not authorized",
	// 	})
	// 	return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"massage": "NOT AUTHORIZED"})
	// 	return
	// }

	var event models.Event

	err := ctx.BindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "could not parse the data",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	// event.ID = 1
	event.UserID = userId

	if err = event.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not create events. try again later",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"massage": "event created",
		"event":   event,
	})

}

func UpdateEvent(ctx *gin.Context) {

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "could not parse event id.",
		})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not fetch the event",
		})
		return
	}

	var UpdatedEvent models.Event
	if err = ctx.BindJSON(&UpdatedEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "could not parse the data",
		})
		return

	}

	UpdatedEvent.ID = eventId
	err = UpdatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not update event.",
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "Event updated successfully."})
}

func DeleteEvent(ctx *gin.Context) {

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
			"massage": "could not fetch the event",
		})
		return
	}
	if err = event.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Could not delete the event.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": " Event Successfully Deleted",
	})

}
