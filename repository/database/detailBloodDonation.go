package database

import (
	"database/sql"
	"mini_project/config"
	"mini_project/models"
)

func GetTotalQtyByPemberianDarah(pemberianDarahID int) (int, error) {
	var totalQty sql.NullInt64
	if err := config.DB.Model(&models.DetailBloodDonation{}).
		Select("COALESCE(SUM(total_qty), 0) AS total_qty").
		Where("pemberian_darah_id = ?", pemberianDarahID).
		Scan(&totalQty).Error; err != nil {
		return 0, err
	}
	if totalQty.Valid {
		return int(totalQty.Int64), nil
	}
	return 0, nil
}

func GetDetailBloodDonation(id uint) (detailDonation *models.DetailBloodDonation, err error) {
	if err = config.DB.First(&detailDonation, id).Error; err != nil {
		return detailDonation, err
	}
	return detailDonation, nil
}
func GetDetailBloodDonationWithTotalQty(id uint) (*models.DetailBloodDonation, int, error) {
	// Mengambil detail donasi darah berdasarkan id
	detailDonation, err := GetDetailBloodDonation(id)
	if err != nil {
		return nil, 0, err
	}

	// Menghitung total_qty berdasarkan pemberian_darah_id
	totalQty, err := GetTotalQtyByPemberianDarah(int(detailDonation.BloodDonationID))
	if err != nil {
		return nil, 0, err
	}

	return detailDonation, totalQty, nil
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
