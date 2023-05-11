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
	donatur := e.Group("/donatur", mid.JWT([]byte(cons.SECRET_JWT)))
	donatur.POST("", controllers.CreateDonaturController)
	donatur.GET("/:id", controllers.GetDonaturIdController)
	donatur.DELETE("/:id", controllers.DeleteDonaturController)
	donatur.PUT("/:id", controllers.UpdateDonaturController)

	//cek riwayat kesehatan
	medicalhistory := e.Group("/medicalhistory", mid.JWT([]byte(cons.SECRET_JWT)))
	medicalhistory.POST("", controllers.CreateMedicalHistoryController)
	medicalhistory.GET("/:id", controllers.GetMedicalHistory)
	medicalhistory.DELETE("/:id", controllers.DeleteMedicalHistoryController)
	medicalhistory.PUT("/:id", controllers.UpdateMedicalHistoryController)

	//cek donasi darah nya
	blooddonation := e.Group("/blooddonation", mid.JWT([]byte(cons.SECRET_JWT)))
	blooddonation.POST("", controllers.CreateBloodDonationController)
	blooddonation.GET("/:id", controllers.GetBloodDonationController)
	blooddonation.DELETE("/:id", controllers.DeleteBloodDonationController)
	blooddonation.PUT("/:id", controllers.UpdateBloodDonationController)

	//cek detail donasi darah nya
	detaildonation := e.Group("/detaildonation", mid.JWT([]byte(cons.SECRET_JWT)))
	detaildonation.POST("", controllers.CreateDetailBloodDonationController)
	// detaildonation.GET("/:id", controllers.GetDetailBloodDonationIdController)
	detaildonation.GET("/:id", controllers.GetDetailBloodDonationTotalQtyController)
	detaildonation.DELETE("/:id", controllers.DeleteDetailBloodDonationController)
	detaildonation.PUT("/:id", controllers.UpdateDetailBloodDonationController)
}
