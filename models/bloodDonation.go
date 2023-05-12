package models

import (
	"gorm.io/gorm"
)

type BloodDonation struct {
	gorm.Model
	DonaturID     uint   `json:"donatur_id" from:"donatur_id" gorm:"foreignkey:DonaturID"`
	PenerimaID    uint   `json:"penerima_id" from:"penerima_id" gorm:"foreignkey:DonaturID"`
	TanggalDonasi string `json:"tanggal_donasi" form:"tanggal_donasi"`
	Qty           string `json:"qty" form:"qty"`
}
