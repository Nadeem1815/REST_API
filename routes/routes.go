package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents) //GET, POST, PATH, DELETE
	server.GET("/events/:id", getEvent)
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
}
