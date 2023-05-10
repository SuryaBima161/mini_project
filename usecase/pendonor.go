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

func CreatePendonor(req *payload.CreatPendonorRequest) (resp payload.CreatPendonorRespone, err error) {

	var existingUser models.Pendonor
	if err := config.DB.First(&existingUser, req.User_id).First(&existingUser).Error; err == nil {
		return payload.CreatPendonorRespone{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "User_id not exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreatPendonorRespone{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	newPendonor := &models.Pendonor{
		User_id:       req.User_id,
		Name:          req.Name,
		Jenis_Kelamin: req.Jenis_Kelamin,
		Tanggal_lahir: req.Tanggal_lahir,
	}
	err = database.CreatePendonor(newPendonor)
	if err != nil {
		return
	}
	resp = payload.CreatPendonorRespone{
		User_id:       newPendonor.User_id,
		Name:          newPendonor.Name,
		Jenis_Kelamin: newPendonor.Jenis_Kelamin,
		Tanggal_lahir: newPendonor.Tanggal_lahir,
	}
	return
}

func GetPendonor(id uint) (*payload.GetPendonorResponse, error) {
	pendonor, err := database.GetPendonor(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetPendonorResponse{
		User_id:       pendonor.User_id,
		Name:          pendonor.Name,
		Jenis_Kelamin: pendonor.Jenis_Kelamin,
		Tanggal_lahir: pendonor.Tanggal_lahir,
	}
	return &resp, nil
}

func UpdatePendonor(pendonor *models.Pendonor) (err error) {
	err = database.UpdatePendonor(pendonor)
	if err != nil {
		fmt.Println("Update : Error updating pendonor, err: ", err)
		return
	}
	return
}
func DeletePendonor(pendonorId uint) (err error) {
	deletePendonor := models.Pendonor{
		ID: pendonorId,
	}
	err = database.DeletePendonor(&deletePendonor)
	if err != nil {
		fmt.Println("Delete: error deleting pendonor, err:", err)
		return
	}
	return err
}
