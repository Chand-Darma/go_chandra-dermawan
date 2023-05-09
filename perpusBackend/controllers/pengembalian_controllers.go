package controllers

import (
	"net/http"
	"project/config"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

//create data pengembalian

func CreatePengembalianController(c echo.Context) error {

	pengembalian := models.Pengembalian{}
	c.Bind(&pengembalian)

	if err := config.DB.Save(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success menambahkan data pengembalian",
		"pengembalian": pengembalian,
	})
}

// get data pengembalian

func GetPengembalianController(c echo.Context) error {
	var pengembalian []models.Pengembalian

	err := config.DB.Find(&pengembalian).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data pengembalian",
		"data":    pengembalian,
	})
}

// get data pengembalian by id

func GetPengembalianIDController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var pengembalian models.Pengembalian
	if err := config.DB.First(&pengembalian, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success menampilkan data pengembalian by id",
		"pengembalian": pengembalian,
	})
}

// update data pengembalian by id

func UpdatePengembalianIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var pengembalian models.Pengembalian

	if err := config.DB.First(&pengembalian, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&pengembalian)

	if err := config.DB.Save(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update data pengembalian by id",
		"pengembalian": pengembalian,
	})
}

func DeletePengembalianIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var pengembalian models.Pengembalian
	if err := config.DB.First(&pengembalian, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pengembalian by id",
	})
}
