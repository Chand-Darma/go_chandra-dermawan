package controllers

import (
	"net/http"
	"project/config"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

//create data peminjam

func CreatePeminjamController(c echo.Context) error {

	peminjam := models.Peminjaman{}
	c.Bind(&peminjam)

	if err := config.DB.Save(&peminjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success menambahkan data peminjam",
		"peminjam": peminjam,
	})
}

// get data peminjaman

func GetPeminjamanController(c echo.Context) error {
	var peminjaman []models.Peminjaman

	err := config.DB.Find(&peminjaman).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menampilkan data peminjaman",
		"data":    peminjaman,
	})
}

// get data peminjam by id

func GetPeminjamanIDController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var peminjam models.Peminjaman
	if err := config.DB.First(&peminjam, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success menampilkan data peminjam by id",
		"peminjam": peminjam,
	})
}

// update data peminjam by id

func UpdatePeminjamIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid admin id")
	}

	var peminjam models.Peminjaman

	if err := config.DB.First(&peminjam, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&peminjam)

	if err := config.DB.Save(&peminjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success update data peminjam by id",
		"peminjam": peminjam,
	})
}

func DeletePeminjamIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var peminjam models.Peminjaman
	if err := config.DB.First(&peminjam, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&peminjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete peminjam by id",
	})
}
