package main

import (
	"eksplorasi/routes"

	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// initialize routes
	routes.InitRoutesBlogs(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
