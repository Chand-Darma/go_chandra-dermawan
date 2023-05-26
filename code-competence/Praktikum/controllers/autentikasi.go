package controllers

import (
	"net/http"

	"project/models"

	"github.com/labstack/echo"
)

var users []models.Users

func Register(c echo.Context) error {
	user := new(models.Users)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	// Check if username already exists
	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			return c.String(http.StatusConflict, "Username already exists")
		}
	}

	// Add new user
	users = append(users, *user)

	return c.String(http.StatusCreated, "User registered successfully")
}

func Login(c echo.Context) error {
	user := new(models.Users)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	// Check if username and password match
	for _, existingUser := range users {
		if existingUser.Username == user.Username && existingUser.Password == user.Password {
			return c.String(http.StatusOK, "Login successful")
		}
	}

	return echo.ErrUnauthorized
}
