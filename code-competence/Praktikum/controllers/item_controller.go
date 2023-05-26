package controllers

import (
	"net/http"
	"strconv"

	"project/config"
	"project/models"

	"github.com/labstack/echo"
)

func GetItemsByCategoryController(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid category ID")
	}

	var filteredItems []models.Items
	if err := config.DB.Where("category_id = ?", categoryID).Find(&filteredItems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    filteredItems,
	})
}

func GetItemsController(c echo.Context) error {
	var items []models.Items

	err := config.DB.Find(&items).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all data",
		"data":    items,
	})
}

func GetItemByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item id")
	}

	var item models.Items
	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by id",
		"item":    item,
	})
}

func CreateItemController(c echo.Context) error {
	item := models.Items{}
	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := config.DB.Save(&item).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create item",
		"item":    item,
	})
}

func UpdateItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item id")
	}

	var item models.Items

	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
		"item":    item,
	})
}

func DeleteItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item id")
	}

	var item models.Items
	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

func GetItemsKeywordController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	var items []models.Items

	if keyword != "" {
		// Jika terdapat keyword, lakukan pencarian berdasarkan nama item
		err := config.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&items).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	} else {
		// Jika keyword kosong, ambil semua item
		err := config.DB.Find(&items).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    items,
	})
}
