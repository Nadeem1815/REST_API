package middlewares

import (
	"net/http"

	"github.com/Nadeem1815/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"massage": "not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"massage": "NOT AUTHORIZED"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
