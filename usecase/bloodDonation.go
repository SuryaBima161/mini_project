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

	var existingBloodDonation models.BloodDonation
	if err := config.DB.Where("donatur_id = ? AND penerima_id = ?", req.DonaturID, req.PenerimaID).First(&existingBloodDonation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return payload.CreateBloodDonationRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message":          "donatur_id or penerima_id does not exist",
				"errorDescription": err.Error(),
			})
		}
		return payload.CreateBloodDonationRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err.Error(),
		})
	}
	newBloodDonation := &models.BloodDonation{
		DonaturID:     req.DonaturID,
		PenerimaID:    req.PenerimaID,
		TanggalDonasi: req.TanggalDonasi,
		Qty:           req.Qty,
	}
	err = database.CreateBloodDonation(newBloodDonation)
	if err != nil {
		return
	}
	resp = payload.CreateBloodDonationRespone{
		DonaturID:     newBloodDonation.DonaturID,
		PenerimaID:    newBloodDonation.PenerimaID,
		TanggalDonasi: newBloodDonation.TanggalDonasi,
		Qty:           newBloodDonation.Qty,
	}
	return
}

func GetBloodDonation(id uint) (*payload.GetBloodDonationResponse, error) {
	GetBloodDonation, err := database.GetBloodDonation(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetBloodDonationResponse{
		DonaturID:     GetBloodDonation.DonaturID,
		PenerimaID:    GetBloodDonation.PenerimaID,
		TanggalDonasi: GetBloodDonation.TanggalDonasi,
		Qty:           GetBloodDonation.Qty,
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
func DeleteBloodDonation(id uint) (err error) {
	DeleteBloodDonation, err := database.GetBloodDonation(id)
	if err != nil {
		return err
	}
	err = database.DeleteBloodDonation(DeleteBloodDonation)
	if err != nil {
		fmt.Println("Delete: error deleting pemberian Darah , err:", err)
		return
	}
	return err
}
