package usecase

import (
	"mini_project/config"
	"mini_project/models"
	"mini_project/models/payload"
	"mini_project/repository/database"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateDonatur(req *payload.CreateDonaturRequest) (resp payload.CreateDonaturRespone, err error) {

	var existingUser models.Donatur
	if err := config.DB.First(&existingUser, req.User_id).First(&existingUser).Error; err == nil {
		return payload.CreateDonaturRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "User_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreateDonaturRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newDonatur := &models.Donatur{
		User_id:       req.User_id,
		Name:          req.Name,
		Jenis_Kelamin: req.Jenis_Kelamin,
		Tanggal_lahir: req.Tanggal_lahir,
	}
	err = database.CreateDonatur(newDonatur)
	if err != nil {
		return
	}
	resp = payload.CreateDonaturRespone{
		User_id:       newDonatur.User_id,
		Name:          newDonatur.Name,
		Jenis_Kelamin: newDonatur.Jenis_Kelamin,
		Tanggal_lahir: newDonatur.Tanggal_lahir,
	}
	return
}

func GetDonatur(id uint) (*payload.GetDonaturResponse, error) {
	Donatur, err := database.GetDonatur(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetDonaturResponse{
		User_id:       Donatur.User_id,
		Name:          Donatur.Name,
		Jenis_Kelamin: Donatur.Jenis_Kelamin,
		Tanggal_lahir: Donatur.Tanggal_lahir,
	}
	return &resp, nil
}

func UpdateDonatur(Donatur *models.Donatur) (err error) {
	err = database.UpdateDonatur(Donatur)
	if err != nil {
		fmt.Println("Update : Error updating Donatur, err: ", err)
		return
	}
	return
}
func DeleteDonatur(DonaturId uint) (err error) {
	deleteDonatur := models.Donatur{
		ID: DonaturId,
	}
	err = database.DeleteDonatur(&deleteDonatur)
	if err != nil {
		fmt.Println("Delete: error deleting Donatur, err:", err)
		return
	}
	return err
}
