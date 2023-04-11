package routes

import (
	"belajar-go-echo/config"
	"belajar-go-echo/constants"
	"belajar-go-echo/controller"
	m "belajar-go-echo/middleware"

	echomid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))

	m.LogMiddlewares(app)

	//jwt auth user
	jwtAuth := app.Group("/jwt")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controller.GetAllUsers(db))
	jwtAuth.POST("/users", controller.CreateUser(db))

	return app
}
