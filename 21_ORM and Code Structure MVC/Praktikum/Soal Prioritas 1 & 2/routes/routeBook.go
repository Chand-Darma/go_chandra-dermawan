package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func InitRoutesBooks(e *echo.Echo) {

	// Route / to handler function
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
