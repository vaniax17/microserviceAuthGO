package user

import (
	"encoding/json"
	"errors"
	"microserviceAuthGO/src/core"
	"microserviceAuthGO/src/db"
	"microserviceAuthGO/src/models"
	"microserviceAuthGO/src/rabbitmq"
	"microserviceAuthGO/src/validators"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	email := c.QueryParam("email")

	valid := validators.IsEmail(email)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email address"})
	}

	valid = validators.IsCorrectUsername(username)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username"})
	}

	PasswordHashSum := core.Hash(password)

	user := &models.User{
		Username:       username,
		Email:          email,
		HashedPassword: PasswordHashSum,
	}
	exist := isExists(user)
	if exist {
		return c.JSON(http.StatusConflict, map[string]string{"message": "User already exists"})
	}

	payload, err := json.Marshal(models.PublishData{Username: username, Email: email})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	msg := message.Message{Payload: payload}
	rabbitmq.PublishUserCreate(&msg)

	db.DB.Create(user)
	token, err := core.CreateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})

}

func isExists(userData *models.User) bool {
	err := db.DB.Model(userData).Where("username = ? OR email = ?", userData.Username, userData.Email).First(userData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
