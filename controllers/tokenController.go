package controllers

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generate token
// @Summary      Generate token
// @Description  Generate token
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param Body body models.Token true "Token Data"
// @Success      200  {object}  models.Token
// @Router       /mysql/token [post]
// @Router 			 /postgresql/token [post]
func GenerateToken(context *gin.Context) {
	var request models.Token
	var client models.Client
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.DB.Where("email = ?", request.Email).First(&client)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := client.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := models.GenerateJWT(client.ID, client.Email, client.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
