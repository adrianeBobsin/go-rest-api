package middlewares

import (
	"jwt-authentication-golang/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		token := strings.Split(tokenString, " ")
		if len(token) != 2 {
			return
		}
		if token[1] == "" {
			context.JSON(401, gin.H{"error": "token was not sended"})
			context.Abort()
			return
		}
		err := models.ValidateToken(token[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
