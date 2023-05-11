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

func CreateDetailBloodDonation(req *payload.CreateDetailBloodDonationRequest) (resp payload.CreateDetailDonationRespone, err error) {

	var existingUser models.BloodDonation
	if err := config.DB.First(&existingUser, req.Pendonor_id, req.Penerima_id).First(&existingUser).Error; err == nil {
		return payload.CreateDetailDonationRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Pendonor_id or Penerima_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreateDetailDonationRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newBloodDonation := &models.DetailBloodDonation{
		Pendonor_id: req.Pendonor_id,
		Penerima_id: req.Penerima_id,
		Riwayat_id:  req.Riwayat_id,
	}
	err = database.CreateDetailBloodDonation(newBloodDonation)
	if err != nil {
		return
	}
	resp = payload.CreateDetailDonationRespone{
		Pendonor_id: newBloodDonation.Pendonor_id,
		Penerima_id: newBloodDonation.Penerima_id,
		Riwayat_id:  newBloodDonation.Riwayat_id,
		Total_Qty:   newBloodDonation.Total_Qty,
	}
	return
}

func GetDetailBloodDonation(id uint) (*payload.GetDetailDonationResponse, error) {
	getDetailDonation, err := database.GetDetailBloodDonation(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetDetailDonationResponse{
		Pendonor_id: getDetailDonation.Pendonor_id,
		Penerima_id: getDetailDonation.Penerima_id,
		Riwayat_id:  getDetailDonation.Riwayat_id,
		Total_Qty:   getDetailDonation.Total_Qty,
	}
	return &resp, nil
}

func GetDetailBloodDonationWithTotalQty(id uint) (*payload.GetDetailDonationResponse, int, error) {
	getDetailDonation, totalQty, err := database.GetDetailBloodDonationWithTotalQty(id)
	if err != nil {
		return nil, 0, err
	}
	resp := payload.GetDetailDonationResponse{
		Pendonor_id: getDetailDonation.Pendonor_id,
		Penerima_id: getDetailDonation.Penerima_id,
		Riwayat_id:  getDetailDonation.Riwayat_id,
		Total_Qty:   getDetailDonation.Total_Qty,
	}
	return &resp, totalQty, nil
}
func UpdateDetailBloodDonation(detailDonation *models.DetailBloodDonation) (err error) {
	err = database.UpdateDetailBloodDonation(detailDonation)
	if err != nil {
		fmt.Println("Update : Error updating detail Donation, err: ", err)
		return
	}
	return
}
func DeleteDetailBloodDonation(detailDonationId uint) (err error) {
	deleteDetailDonation := models.DetailBloodDonation{
		ID: detailDonationId,
	}
	err = database.DeleteDetailBloodDonation(&deleteDetailDonation)
	if err != nil {
		fmt.Println("Delete: error deleting detail donation , err:", err)
		return
	}
	return err
}
