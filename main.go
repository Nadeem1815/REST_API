package main

import (
	"github.com/Nadeem1815/rest-api/db"
	"github.com/Nadeem1815/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
