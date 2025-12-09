package routes

import (
	"net/http"

	"github.com/Nadeem1815/rest-api/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "Request filed",
		})
		return
	}

	if err := user.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "could not save user, try again ",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"massage": "user created",
		"event":   user,
	})

}
