package main

import (
	"mini_project/config"
	"mini_project/middleware"
	routes "mini_project/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	middleware.LogMiddleware(e)
	db := config.InitDB()
	routes.New(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
