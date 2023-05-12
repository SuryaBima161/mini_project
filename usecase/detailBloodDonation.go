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

	var existingBloodDonation models.BloodDonation
	if err := config.DB.Where("donatur_id = ? AND penerima_id = ?", req.DonaturID, req.PenerimaID).First(&existingBloodDonation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return payload.CreateDetailDonationRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message":          "donatur_id or penerima_id does not exist",
				"errorDescription": err.Error(),
			})
		}
		return payload.CreateDetailDonationRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err.Error(),
		})
	}
	detailDonation := &models.DetailBloodDonation{
		DonaturID:        req.DonaturID,
		PenerimaID:       req.PenerimaID,
		MedicalHistoryID: req.MedicalHistoryID,
		BloodDonationID:  req.BloodDonationID,
	}
	err = database.CreateDetailBloodDonation(detailDonation)
	if err != nil {
		return
	}
	resp = payload.CreateDetailDonationRespone{
		DonaturID:        detailDonation.DonaturID,
		PenerimaID:       detailDonation.PenerimaID,
		MedicalHistoryID: detailDonation.MedicalHistoryID,
		BloodDonationID:  req.BloodDonationID,
	}
	return
}

func GetDetailBloodDonation(id uint) (*payload.GetDetailDonationResponse, error) {
	getDetailDonation, err := database.GetDetailBloodDonation(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetDetailDonationResponse{
		TotalQty: getDetailDonation.TotalQty,
	}
	return &resp, nil
}

func GetDetailBloodDonationWithTotalQty(id uint) (*payload.GetDetailDonationResponse, int, error) {
	getDetailDonation, totalQty, err := database.GetDetailBloodDonationWithTotalQty(id)
	if err != nil {
		return nil, 0, err
	}
	resp := payload.GetDetailDonationResponse{
		TotalQty: getDetailDonation.TotalQty,
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
