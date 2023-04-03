package main

import (
	"project/routes"

	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// initialize routes
	routes.InitRoutesUsers(e)
	routes.InitRoutesBooks(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
