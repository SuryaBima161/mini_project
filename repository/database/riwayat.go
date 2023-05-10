package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetRiwayat(id uint) (riwayat *models.Riwayat, err error) {
	if err = config.DB.First(&riwayat, id).Error; err != nil {
		return riwayat, err
	}
	return riwayat, nil
}

func CreateRiwayat(riwayat *models.Riwayat) (err error) {

	if err = config.DB.Create(&riwayat).Error; err != nil {
		return err
	}
	return nil
}
func UpdateRiwayat(riwayat *models.Riwayat) error {
	if err := config.DB.Updates(riwayat).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRiwayat(riwayat *models.Riwayat) error {
	if err := config.DB.Delete(riwayat).Error; err != nil {
		return err
	}
	return nil
}
