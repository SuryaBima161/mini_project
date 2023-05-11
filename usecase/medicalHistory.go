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
	if err := config.DB.First(&existingUser, req.Pendonor_id).First(&existingUser).Error; err == nil {
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
		Pendonor_id:         req.Pendonor_id,
		Tanggal_Pemeriksaan: req.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   req.Hasil_Pemeriksaan,
		Golongan_Darah:      req.Golongan_Darah,
	}
	err = database.CreateMedicalHistory(newMedicalHistory)
	if err != nil {
		return
	}
	resp = payload.CreatMedicalHistoryRespone{
		Pendonor_id:         newMedicalHistory.Pendonor_id,
		Tanggal_Pemeriksaan: newMedicalHistory.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   newMedicalHistory.Hasil_Pemeriksaan,
		Golongan_Darah:      newMedicalHistory.Golongan_Darah,
	}
	return
}

func GetMedicalHistory(id uint) (*payload.GetMedicalHistoryResponse, error) {
	getMedicalHistory, err := database.GetMedicalHistory(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetMedicalHistoryResponse{
		Pendonor_id:         getMedicalHistory.Pendonor_id,
		Tanggal_Pemeriksaan: getMedicalHistory.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   getMedicalHistory.Hasil_Pemeriksaan,
		Golongan_Darah:      getMedicalHistory.Golongan_Darah,
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
func DeleteMedicalHistory(MedicalHistoryId uint) (err error) {
	deleteMedicalHistory := models.MedicalHistory{
		ID: MedicalHistoryId,
	}
	err = database.DeleteMedicalHistory(&deleteMedicalHistory)
	if err != nil {
		fmt.Println("Delete: error deleting MedicalHistory, err:", err)
		return
	}
	return err
}
