package middlewares

import (
	"jwt-authentication-golang/database"
	"strings"

	"github.com/gin-gonic/gin"
)

func SelectBD() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.URL.Path
		token := strings.Split(tokenString, "/")
		if token[3] == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		if token[3] == "mysql" {
			database.ConnectionDBMySQL()
			database.Migrate()
		}

		if token[3] == "postgresql" {
			database.ConnectionDBPostegreSQL()
			database.Migrate()
		}
		context.Next()
	}
}
