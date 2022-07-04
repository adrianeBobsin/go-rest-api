package main

import (
	"jwt-authentication-golang/docs"
	"jwt-authentication-golang/routes"
)

func main() {
	// Swagger header
	docs.SwaggerInfo.Title = "Go rest API"
	docs.SwaggerInfo.Description = "API build in Golang for insert clients contacts in DB."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := routes.InitRouter()
	router.Run(":8080")
}
