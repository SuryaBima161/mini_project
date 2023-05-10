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

func CreateRiwayat(req *payload.CreateRiwayatRequest) (resp payload.CreatRiwayatRespone, err error) {

	var existingUser models.Riwayat
	if err := config.DB.First(&existingUser, req.Pendonor_id).First(&existingUser).Error; err == nil {
		return payload.CreatRiwayatRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Pendonor_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreatRiwayatRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newRiwayat := &models.Riwayat{
		Pendonor_id:         req.Pendonor_id,
		Tanggal_Pemeriksaan: req.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   req.Hasil_Pemeriksaan,
		Golongan_Darah:      req.Golongan_Darah,
	}
	err = database.CreateRiwayat(newRiwayat)
	if err != nil {
		return
	}
	resp = payload.CreatRiwayatRespone{
		Pendonor_id:         newRiwayat.Pendonor_id,
		Tanggal_Pemeriksaan: newRiwayat.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   newRiwayat.Hasil_Pemeriksaan,
		Golongan_Darah:      newRiwayat.Golongan_Darah,
	}
	return
}

func GetRiwayat(id uint) (*payload.GetRiwayatResponse, error) {
	getRiwayat, err := database.GetRiwayat(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetRiwayatResponse{
		Pendonor_id:         getRiwayat.Pendonor_id,
		Tanggal_Pemeriksaan: getRiwayat.Tanggal_Pemeriksaan,
		Hasil_Pemeriksaan:   getRiwayat.Hasil_Pemeriksaan,
		Golongan_Darah:      getRiwayat.Golongan_Darah,
	}
	return &resp, nil
}

func UpdateRiwayat(riwayat *models.Riwayat) (err error) {
	err = database.UpdateRiwayat(riwayat)
	if err != nil {
		fmt.Println("Update : Error updating Riwayat, err: ", err)
		return
	}
	return
}
func DeleteRiwayat(riwayatId uint) (err error) {
	deleteRiwayat := models.Riwayat{
		ID: riwayatId,
	}
	err = database.DeleteRiwayat(&deleteRiwayat)
	if err != nil {
		fmt.Println("Delete: error deleting Riwayat, err:", err)
		return
	}
	return err
}
