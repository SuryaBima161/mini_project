package routes

import (
	cons "mini_project/constant"
	"mini_project/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//login n sign in
	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)

	//JWT
	user := e.Group("", middleware.JWT([]byte(cons.SECRET_JWT)))
	user.GET("/users", controllers.GetUserControllers)
	user.GET("/users/:id", controllers.GetUsersIdController)
	user.DELETE("/users/:id", controllers.DeleteUserController)
	user.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
