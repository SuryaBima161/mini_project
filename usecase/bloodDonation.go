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

func CreateBloodDonation(req *payload.CreateBloodDonationRequest) (resp payload.CreateBloodDonationRespone, err error) {

	var existingUser models.BloodDonation
	if err := config.DB.First(&existingUser, req.Pendonor_id, req.Penerima_id).First(&existingUser).Error; err == nil {
		return payload.CreateBloodDonationRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Pendonor_id or Penerima_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreateBloodDonationRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newBloodDonation := &models.BloodDonation{
		Pendonor_id:    req.Pendonor_id,
		Penerima_id:    req.Penerima_id,
		Tanggal_Donasi: req.Tanggal_Donasi,
		Qty:            req.Qty,
	}
	err = database.CreateBloodDonation(newBloodDonation)
	if err != nil {
		return
	}
	resp = payload.CreateBloodDonationRespone{
		Pendonor_id:    newBloodDonation.Pendonor_id,
		Penerima_id:    newBloodDonation.Penerima_id,
		Tanggal_Donasi: newBloodDonation.Tanggal_Donasi,
		Qty:            newBloodDonation.Qty,
	}
	return
}

func GetBloodDonation(id uint) (*payload.GetBloodDonationResponse, error) {
	GetBloodDonation, err := database.GetBloodDonation(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetBloodDonationResponse{
		Pendonor_id:    GetBloodDonation.Pendonor_id,
		Penerima_id:    GetBloodDonation.Penerima_id,
		Tanggal_Donasi: GetBloodDonation.Tanggal_Donasi,
		Qty:            GetBloodDonation.Qty,
	}
	return &resp, nil
}

func UpdateBloodDonation(BloodDonation *models.BloodDonation) (err error) {
	err = database.UpdateBloodDonation(BloodDonation)
	if err != nil {
		fmt.Println("Update : Error updating Pemberian Darah, err: ", err)
		return
	}
	return
}
func DeleteBloodDonation(BloodDonationId uint) (err error) {
	DeleteBloodDonation := models.BloodDonation{
		ID: BloodDonationId,
	}
	err = database.DeleteBloodDonation(&DeleteBloodDonation)
	if err != nil {
		fmt.Println("Delete: error deleting pemberian Darah , err:", err)
		return
	}
	return err
}
