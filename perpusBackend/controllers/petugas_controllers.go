package controllers

import (
	"net/http"
	"strconv"

	"project/config"
	"project/models"

	"github.com/labstack/echo"
)

// create data petugas

func CreatePetugasController(c echo.Context) error {
	petugas := models.Petugas{}
	c.Bind(&petugas)

	if err := config.DB.Save(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new data petugas",
		"book":    petugas,
	})
}

// get data petugas
func GetPetugasController(c echo.Context) error {
	var petugas []models.Petugas

	err := config.DB.Find(&petugas).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data petugas",
		"data":    petugas,
	})
}

// get petugas by id
func GetPetugasIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var petugas models.Petugas
	if err := config.DB.First(&petugas, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data petugas by id",
		"petugas": petugas,
	})
}

// update data petugas by id

func UpdatePetugasIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var petugas models.Petugas

	if err := config.DB.First(&petugas, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&petugas)

	if err := config.DB.Save(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update petugas by id",
		"petugas": petugas,
	})
}

// delete petugas by id
func DeletePetugasIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var petugas models.Petugas
	if err := config.DB.First(&petugas, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete petugas by id",
	})
}
