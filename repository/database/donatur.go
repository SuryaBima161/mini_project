package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetDonatur(id uint) (donatur models.Donatur, err error) {
	// if err = config.DB.Preload("DetailBloodDonation").Find(&donatur, id).Error; err != nil {
	// 	return donatur, err
	// }
	// return donatur, nil
	if err = config.DB.Find(&donatur, id).Error; err != nil {
		return donatur, err
	}
	return donatur, nil
}

func CreateDonatur(donatur *models.Donatur) (err error) {

	if err = config.DB.Create(&donatur).Error; err != nil {
		return err
	}
	return nil
}
func UpdateDonatur(donatur *models.Donatur) error {
	if err := config.DB.Updates(donatur).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDonatur(donatur *models.Donatur) error {
	if err := config.DB.Delete(donatur).Error; err != nil {
		return err
	}
	return nil
}
