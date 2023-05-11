package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetBloodDonation(id uint) (donation *models.BloodDonation, err error) {
	if err = config.DB.First(&donation, id).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func CreateBloodDonation(donation *models.BloodDonation) (err error) {

	if err = config.DB.Create(&donation).Error; err != nil {
		return err
	}
	return nil
}
func UpdateBloodDonation(donation *models.BloodDonation) error {
	if err := config.DB.Updates(donation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBloodDonation(donation *models.BloodDonation) error {
	if err := config.DB.Delete(donation).Error; err != nil {
		return err
	}
	return nil
}
