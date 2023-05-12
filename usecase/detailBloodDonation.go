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
	if err := config.DB.First(&existingUser, req.DonaturID, req.PenerimaID).First(&existingUser).Error; err == nil {
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
	detailDonation := &models.DetailBloodDonation{
		DonaturID:        req.DonaturID,
		PenerimaID:       req.PenerimaID,
		MedicalHistoryID: req.MedicalHistoryID,
	}
	err = database.CreateDetailBloodDonation(detailDonation)
	if err != nil {
		return
	}
	resp = payload.CreateDetailDonationRespone{
		DonaturID:        detailDonation.DonaturID,
		PenerimaID:       detailDonation.PenerimaID,
		MedicalHistoryID: detailDonation.MedicalHistoryID,
		TotalQty:         detailDonation.TotalQty,
	}
	return
}

func GetDetailBloodDonation(id uint) (*payload.GetDetailDonationResponse, error) {
	getDetailDonation, err := database.GetDetailBloodDonation(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetDetailDonationResponse{
		DonaturID:        getDetailDonation.DonaturID,
		PenerimaID:       getDetailDonation.PenerimaID,
		MedicalHistoryID: getDetailDonation.MedicalHistoryID,
		TotalQty:         getDetailDonation.TotalQty,
	}
	return &resp, nil
}

func GetDetailBloodDonationWithTotalQty(id uint) (*payload.GetDetailDonationResponse, int, error) {
	getDetailDonation, totalQty, err := database.GetDetailBloodDonationWithTotalQty(id)
	if err != nil {
		return nil, 0, err
	}
	resp := payload.GetDetailDonationResponse{
		DonaturID:        getDetailDonation.DonaturID,
		PenerimaID:       getDetailDonation.PenerimaID,
		MedicalHistoryID: getDetailDonation.MedicalHistoryID,
		TotalQty:         getDetailDonation.TotalQty,
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
func DeleteDetailBloodDonation(id uint) (err error) {
	deleteDetailDonation := models.DetailBloodDonation{
		Model: gorm.Model{ID: id},
	}
	err = database.DeleteDetailBloodDonation(&deleteDetailDonation)
	if err != nil {
		fmt.Println("Delete: error deleting detail donation , err:", err)
		return
	}
	return err
}
