package controllers

import (
	"net/http"
	"project/config"
	"project/middleware"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// get all admin

func GetAdminsController(c echo.Context) error {
	var admins []models.Admins

	err := config.DB.Find(&admins).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data admin",
		"data":    admins,
	})
}

// get admin by id
func GetAdminsIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var admin models.Admins
	if err := config.DB.First(&admin, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data admin by id",
		"admin":   admin,
	})
}

// create admin
func CreateAdminController(c echo.Context) error {

	admin := models.Admins{}
	c.Bind(&admin)

	err := config.DB.Save(&admin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(admin.ID), admin.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	admin.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create admin",
		"admin":   admin,
	})
}

// update data admin by id

func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var admins models.Admins

	if err := config.DB.First(&admins, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&admins)

	if err := config.DB.Save(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin",
		"user":    admins,
	})
}

// delete data admin

func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var admin models.Admins
	if err := config.DB.First(&admin, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete admin",
	})
}
