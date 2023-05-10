package routes

import (
	cons "mini_project/constant"
	"mini_project/controllers"
	"mini_project/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	//login n register
	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)

	//user account
	user := e.Group("/users", mid.JWT([]byte(cons.SECRET_JWT)))
	user.GET("/:id", controllers.GetUsersIdController)
	user.DELETE("/:id", controllers.DeleteUserController)
	user.PUT("/:id", controllers.UpdateUserController)

	//waktu donor
	pendonor := e.Group("/pendonor", mid.JWT([]byte(cons.SECRET_JWT)))
	pendonor.POST("", controllers.CreatePendonorController)
	pendonor.GET("/:id", controllers.GetPendonorId)
	pendonor.DELETE("/:id", controllers.DeletePendonorController)
	pendonor.PUT("/:id", controllers.UpdatePendonorController)
}
