package routes

import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	mysql := router.Group("/api/v1/mysql/").Use(middlewares.SelectBD())
	{
		mysql.POST("token", controllers.GenerateToken)
		mysql.POST("client/create", controllers.CreateClient)
		mysql.Use(middlewares.Auth()).POST("client/contacts", controllers.CreateContact)
	}

	postgresql := router.Group("api/v1/postgresql/").Use(middlewares.SelectBD())
	{
		postgresql.POST("token", controllers.GenerateToken)
		postgresql.POST("client/create", controllers.CreateClient)
		postgresql.Use(middlewares.Auth()).POST("client/contacts", controllers.CreateContact)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
