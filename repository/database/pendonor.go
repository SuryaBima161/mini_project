package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetPendonor(id uint) (pendonor *models.Pendonor, err error) {
	if err = config.DB.First(&pendonor, id).Error; err != nil {
		return pendonor, err
	}
	return pendonor, nil
}

func CreatePendonor(pendonor *models.Pendonor) (err error) {

	if err = config.DB.Create(&pendonor).Error; err != nil {
		return err
	}
	return nil
}
func UpdatePendonor(pendonor *models.Pendonor) error {
	if err := config.DB.Updates(pendonor).Error; err != nil {
		return err
	}
	return nil
}

func DeletePendonor(pendonor *models.Pendonor) error {
	if err := config.DB.Delete(pendonor).Error; err != nil {
		return err
	}
	return nil
}
