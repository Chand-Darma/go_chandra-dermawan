package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func InitRoutesUsers(e *echo.Echo) {

	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

}
