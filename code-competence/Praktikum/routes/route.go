package routes

import (
	"project/constants"
	"project/controllers"
	m "project/middleware"

	echomid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//jwt auth items
	jwtAuth := e.Group("/jwt")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/items", controllers.GetItemsController)
	jwtAuth.GET("/items/:id", controllers.GetItemByIDController)
	jwtAuth.POST("/items", controllers.CreateItemController)
	jwtAuth.PUT("/items/:id", controllers.UpdateItemController)
	jwtAuth.DELETE("/items/:id", controllers.DeleteItemController)
	jwtAuth.GET("/items/category/:category_id", controllers.GetItemsByCategoryController)
	jwtAuth.GET("/items?keyword=item_name", controllers.GetItemsKeywordController)
	jwtAuth.POST("/login", controllers.Login)

	m.LogMiddlewares(e)

	// autentikasi
	e.POST("/register", controllers.Register)

	return e
}
