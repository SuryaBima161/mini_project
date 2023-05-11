package database

import (
	"mini_project/config"
	"mini_project/models"
)

func GetMedicalHistory(id uint) (medicalHistory *models.MedicalHistory, err error) {
	if err = config.DB.First(&medicalHistory, id).Error; err != nil {
		return medicalHistory, err
	}
	return medicalHistory, nil
}

func CreateMedicalHistory(medicalHistory *models.MedicalHistory) (err error) {

	if err = config.DB.Create(&medicalHistory).Error; err != nil {
		return err
	}
	return nil
}
func UpdateMedicalHistory(medicalHistory *models.MedicalHistory) error {
	if err := config.DB.Updates(medicalHistory).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMedicalHistory(medicalHistory *models.MedicalHistory) error {
	if err := config.DB.Delete(medicalHistory).Error; err != nil {
		return err
	}
	return nil
}
