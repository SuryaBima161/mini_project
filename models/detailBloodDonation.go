package models

import (
	"gorm.io/gorm"
)

type DetailBloodDonation struct {
	gorm.Model
	DonaturID        uint   `json:"donatur_id" from:"donatur_id" gorm:"foreignkey:DonaturID"`
	PenerimaID       uint   `json:"penerima_id" from:"penerima_id" gorm:"foreignkey:DonaturID"`
	MedicalHistoryID uint   `json:"medical_history_id" form:"medical_history_id" gorm:"foreignkey:MedicalHistoryID"`
	BloodDonationID  uint   `json:"blood_donation_id" form:"blood_donation_id" gorm:"foreignkey:BloodDonationID"`
	TotalQty         string `json:"total_qty" form:"total_qty"`
}
