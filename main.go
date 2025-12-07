package main

import (
	"net/http"

	"github.com/Nadeem1815/rest-api/db"
	"github.com/Nadeem1815/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", GetEvents) //GET, POST, PATH, DELETE
	server.POST("/events", CreateEvent)

	server.Run(":8080") // localhost:8080
}

func GetEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.BindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "could not parse the data",
		})
		return
	}

	// event.ID = 1
	// event.UserID = 1

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"massage": "event created",
		"event":   event,
	})

}
