package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (client *Client) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	client.Password = string(bytes)
	return nil
}

func (client *Client) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
