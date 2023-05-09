package controllers

import (
	"net/http"
	"strconv"

	"project/config"
	"project/models"

	"github.com/labstack/echo"
)

// get all data anggota
func GetAnggotaController(c echo.Context) error {
	var anggota []models.Anggota

	err := config.DB.Find(&anggota).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all data anggota",
		"data":    anggota,
	})
}

// get data anggota by id
func GetAnggotaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var anggota models.Anggota
	if err := config.DB.First(&anggota, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data anggota by id",
		"anggota": anggota,
	})
}

//create data anggota

func CreateAnggotaController(c echo.Context) error {

	anggota := models.Anggota{}
	c.Bind(&anggota)

	if err := config.DB.Save(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new data anggota",
		"anggota": anggota,
	})
}

func UpdateAnggotaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var anggota models.Anggota

	if err := config.DB.First(&anggota, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&anggota)

	if err := config.DB.Save(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data anggota",
		"anggota": anggota,
	})
}

// delete user by id
func DeleteAnggotaIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var anggota models.Anggota
	if err := config.DB.First(&anggota, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete anggota",
	})
}
