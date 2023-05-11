package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetDetailBloodDonation(id uint) (detailDonation *models.DetailBloodDonation, err error) {
	if err = config.DB.First(&detailDonation, id).Error; err != nil {
		return detailDonation, err
	}
	return detailDonation, nil
}

func CreateDetailBloodDonation(detailDonation *models.DetailBloodDonation) (err error) {

	if err = config.DB.Create(&detailDonation).Error; err != nil {
		return err
	}
	return nil
}
func UpdateDetailBloodDonation(detailDonation *models.DetailBloodDonation) error {
	if err := config.DB.Updates(detailDonation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDetailBloodDonation(detailDonation *models.DetailBloodDonation) error {
	if err := config.DB.Delete(detailDonation).Error; err != nil {
		return err
	}
	return nil
}
