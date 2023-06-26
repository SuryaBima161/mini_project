package main

import (
	"mini_project/config"
	routes "mini_project/router"
)

func init() {
	config.InitDB()
	config.InitMigrate()
}

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
