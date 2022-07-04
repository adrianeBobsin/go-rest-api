package controllers

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create client
// @Summary      Create client
// @Description  Create clients
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param Body body models.Client true "Client Data"
// @Success      200  {object}  models.Client
// @Router       /mysql/client/create [post]
// @Router 			 /postgresql/client/create [post]
func CreateClient(context *gin.Context) {
	var client models.Client
	if err := context.ShouldBindJSON(&client); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := client.HashPassword(client.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.DB.Create(&client)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"clientId": client.ID, "email": client.Email, "clientname": client.Name})
}
