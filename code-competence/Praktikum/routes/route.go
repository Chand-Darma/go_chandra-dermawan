package routes

import (
	"project/controllers"
	m "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// items
	e.GET("/items", controllers.GetItemsController)
	e.GET("/items/:id", controllers.GetItemByIDController)
	e.POST("/items", controllers.CreateItemController)
	//e.POST("/users/login", controllers.LoginUsersController)
	e.DELETE("/items/:id", controllers.DeleteItemController)
	//e.PUT("/users/:id", controllers.UpdateUserController)
	e.GET("/items/category/:category_id", controllers.GetItemsByCategoryController)
	e.GET("/items?keyword=item_name", controllers.GetItemsKeywordController)
	e.PUT("/items/:id", controllers.UpdateItemController)

	m.LogMiddlewares(e)

	// autentikasi
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	//auth
	//eBasicAuth := e.Group("/auth")
	//eBasicAuth.Use(echomid.BasicAuth(m.BasicAuthDB))
	//eBasicAuth.GET("/users", controllers.GetUserController)

	//jwt auth user
	//jwtAuth := e.Group("/jwt")
	//jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	//jwtAuth.GET("/users", controllers.GetUsersController)
	//jwtAuth.GET("/users/:id", controllers.GetUserController)
	//jwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	//jwtAuth.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
