package routes

import (
	"github.com/Nadeem1815/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents) //GET, POST, PATH, DELETE
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}
