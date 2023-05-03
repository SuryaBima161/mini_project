package main

import (
	"mini_project/config"
	"mini_project/middleware"
	routes "mini_project/router"
)

func main() {
	config.InitDB()
	e := routes.New()
	middleware.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8080"))
}
