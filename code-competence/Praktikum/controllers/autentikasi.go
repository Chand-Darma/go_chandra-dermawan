package controllers

import (
	"net/http"

	"project/config"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
)

var users []models.Users

func Register(c echo.Context) error {

	user := models.Users{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success register",
		"user":    user,
	})
}

func Login(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	var existingUser models.Users
	err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error
	if err != nil {
		return echo.ErrUnauthorized
	}

	if user.Password != existingUser.Password {
		return echo.ErrUnauthorized
	}

	token, err := middleware.CreateToken(int(existingUser.ID), existingUser.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create JWT Token",
			"status":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
