package routes

import (
	"project/constants"
	"project/controllers"
	m "project/middleware"

	"github.com/labstack/echo"
	echomid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//user route
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUsersController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	m.LogMiddlewares(e)

	// books route
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	//auth
	//eBasicAuth := e.Group("/auth")
	//eBasicAuth.Use(echomid.BasicAuth(m.BasicAuthDB))
	//eBasicAuth.GET("/users", controllers.GetUserController)

	//jwt auth user
	jwtAuth := e.Group("/jwt")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controllers.GetUsersController)
	jwtAuth.GET("/users/:id", controllers.GetUserController)
	jwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	jwtAuth.PUT("/users/:id", controllers.UpdateUserController)

	//jwt auth book
	jwtAuth.GET("/books", controllers.GetBooksController)
	jwtAuth.GET("/books/:id", controllers.GetBookController)
	jwtAuth.POST("/books", controllers.CreateBookController)
	jwtAuth.DELETE("/books/:id", controllers.DeleteBookController)
	jwtAuth.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
