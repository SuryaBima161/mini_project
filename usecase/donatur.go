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
	if err := config.DB.First(&existingUser, req.UserID).First(&existingUser).Error; err == nil {
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
		UserID:       req.UserID,
		Name:         req.Name,
		JenisKelamin: req.JenisKelamin,
		TanggalLahir: req.Tanggallahir,
	}
	err = database.CreateDonatur(newDonatur)
	if err != nil {
		return
	}
	resp = payload.CreateDonaturRespone{
		UserID:       newDonatur.UserID,
		Name:         newDonatur.Name,
		JenisKelamin: newDonatur.JenisKelamin,
		Tanggallahir: newDonatur.TanggalLahir,
	}
	return
}

func GetDonatur(id uint) (*payload.GetDonaturResponse, error) {
	Donatur, err := database.GetDonatur(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetDonaturResponse{
		UserID:       Donatur.UserID,
		Name:         Donatur.Name,
		JenisKelamin: Donatur.JenisKelamin,
		Tanggallahir: Donatur.TanggalLahir,
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
func DeleteDonatur(id uint) (err error) {
	deleteDonatur := models.Donatur{
		Model: gorm.Model{ID: id},
	}
	err = database.DeleteDonatur(&deleteDonatur)
	if err != nil {
		fmt.Println("Delete: error deleting Donatur, err:", err)
		return
	}
	return err
}
