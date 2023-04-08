package controllers

import (
	"net/http"
	"strconv"

	"project/config"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
)

// get all book
func GetBooksController(c echo.Context) error {
	var books []models.Books

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Books{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func noneCreateBookController(c echo.Context) error {

	book := models.Books{}
	c.Bind(&book)

	err := config.DB.Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(book.ID), book.Judul)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	book.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var book models.Books

	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
