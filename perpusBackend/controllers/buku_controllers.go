package controllers

import (
	"net/http"
	"project/config"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// create admin
func CreateBukuController(c echo.Context) error {

	buku := models.Buku{}
	c.Bind(&buku)

	err := config.DB.Save(&buku).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data buku",
		"buku":    buku,
	})
}

// get all buku

func GetBukuController(c echo.Context) error {
	var buku []models.Buku

	err := config.DB.Find(&buku).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data semua buku",
		"data":    buku,
	})
}

// get data buku by id
func GetBukuIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid buku id")
	}

	var buku models.Buku
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data buku by id",
		"anggota": buku,
	})
}

// update data buku by id
func UpdateBukuIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var buku models.Buku

	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&buku)

	if err := config.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update buku by id",
		"buku":    buku,
	})
}

// delete user by id
func DeleteBukuIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var buku models.Buku
	if err := config.DB.First(&buku, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menghapus data buku",
	})
}
