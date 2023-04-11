package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/routes"
)

func main() {
	config.ConnectDB()
	app := routes.New()
	app.Start(":8080")
}
