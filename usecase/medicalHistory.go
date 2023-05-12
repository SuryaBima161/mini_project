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

func CreateMedicalHistory(req *payload.CreateMedicalHistoryRequest) (resp payload.CreatMedicalHistoryRespone, err error) {

	var existingUser models.MedicalHistory
	if err := config.DB.First(&existingUser, req.DonaturID).First(&existingUser).Error; err == nil {
		return payload.CreatMedicalHistoryRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Pendonor_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreatMedicalHistoryRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newMedicalHistory := &models.MedicalHistory{
		DonaturID:          req.DonaturID,
		TanggalPemeriksaan: req.TanggalPemeriksaan,
		HasilPemeriksaan:   req.HasilPemeriksaan,
		GolonganDarah:      req.GolonganDarah,
	}
	err = database.CreateMedicalHistory(newMedicalHistory)
	if err != nil {
		return
	}
	resp = payload.CreatMedicalHistoryRespone{
		DonaturID:          newMedicalHistory.DonaturID,
		TanggalPemeriksaan: newMedicalHistory.TanggalPemeriksaan,
		HasilPemeriksaan:   newMedicalHistory.HasilPemeriksaan,
		GolonganDarah:      newMedicalHistory.GolonganDarah,
	}
	return
}

func GetMedicalHistory(id uint) (*payload.GetMedicalHistoryResponse, error) {
	getMedicalHistory, err := database.GetMedicalHistory(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetMedicalHistoryResponse{
		DonaturID:          getMedicalHistory.DonaturID,
		TanggalPemeriksaan: getMedicalHistory.TanggalPemeriksaan,
		HasilPemeriksaan:   getMedicalHistory.HasilPemeriksaan,
		GolonganDarah:      getMedicalHistory.GolonganDarah,
	}
	return &resp, nil
}

func UpdateMedicalHistory(MedicalHistory *models.MedicalHistory) (err error) {
	err = database.UpdateMedicalHistory(MedicalHistory)
	if err != nil {
		fmt.Println("Update : Error updating MedicalHistory, err: ", err)
		return
	}
	return
}
func DeleteMedicalHistory(id uint) (err error) {
	deleteMedicalHistory := models.MedicalHistory{
		Model: gorm.Model{ID: id},
	}
	err = database.DeleteMedicalHistory(&deleteMedicalHistory)
	if err != nil {
		fmt.Println("Delete: error deleting MedicalHistory, err:", err)
		return
	}
	return err
}
