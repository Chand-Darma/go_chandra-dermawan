package routes

import (
	"eksplorasi/controllers"

	"github.com/labstack/echo"
)

func InitRoutesBlogs(e *echo.Echo) {

	// Route / to handler function
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
}
