package controllers

import (
	"encoding/json"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create contacts
// @Summary      Create contacts
// @Description  Create many clients contacts with one token
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param Body body models.ContactsList true "Contact Data"
// @securityDefinitions Bearer JWT
// @Success      200  {object}  models.Contacts
// @Router       /mysql/client/contacts [post]
// @Router 			 /postgresql/client/contacts [post]
func CreateContact(context *gin.Context) {
	var newContact models.Contacts
	if err := context.ShouldBindJSON(&newContact); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	for i := range newContact.Contacts {
		newContact.Contacts[i].Id_cliente = models.IdToken
	}
	database.DB.Table("contacts").Create(&newContact.Contacts)
	json.NewEncoder(context.Writer).Encode(newContact)
}
