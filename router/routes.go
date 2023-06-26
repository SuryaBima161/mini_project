package routes

import (
	"mini_project/controllers"
	midd "mini_project/middleware"
	"mini_project/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.Validator = &util.CustomValidator{Validator: validator.New()}
	midd.LogMiddleware(e)

	//login n register
	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)

	//user account
	user := e.Group("/users", midd.Cookie)
	user.GET("/:id", controllers.GetUsersIdController)
	user.DELETE("/:id", controllers.DeleteUserController)
	user.PUT("/:id", controllers.UpdateUserController)

	//waktu donor
	donatur := e.Group("/donatur", midd.Cookie)
	donatur.POST("", controllers.CreateDonaturController)
	donatur.GET("/:id", controllers.GetDonaturIdController)
	donatur.DELETE("/:id", controllers.DeleteDonaturController)
	donatur.PUT("/:id", controllers.UpdateDonaturController)

	//cek riwayat kesehatan
	medicalhistory := e.Group("/medicalhistory", midd.Cookie)
	medicalhistory.POST("", controllers.CreateMedicalHistoryController)
	medicalhistory.GET("/:id", controllers.GetMedicalHistory)
	medicalhistory.DELETE("/:id", controllers.DeleteMedicalHistoryController)
	medicalhistory.PUT("/:id", controllers.UpdateMedicalHistoryController)

	//cek donasi darah nya
	blooddonation := e.Group("/blooddonation", midd.Cookie)
	blooddonation.POST("", controllers.CreateBloodDonationController)
	blooddonation.GET("/:id", controllers.GetBloodDonationController)
	blooddonation.DELETE("/:id", controllers.DeleteBloodDonationController)
	blooddonation.PUT("/:id", controllers.UpdateBloodDonationController)

	//cek detail donasi darah nya
	detaildonation := e.Group("/detaildonation", midd.Cookie)
	detaildonation.POST("", controllers.CreateDetailBloodDonationController)
	// detaildonation.GET("/:id", controllers.GetDetailBloodDonationIdController)
	detaildonation.GET("/:id", controllers.GetDetailBloodDonationTotalQtyController)
	detaildonation.DELETE("/:id", controllers.DeleteDetailBloodDonationController)
	detaildonation.PUT("/:id", controllers.UpdateDetailBloodDonationController)

	return e
}
