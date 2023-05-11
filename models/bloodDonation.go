package models

import (
	"gorm.io/gorm"
)

type BloodDonation struct {
	gorm.Model
	ID             uint   `json:"id" from:"id"`
	Pendonor_id    uint   `json:"pendonor_id" from:"pendonor_id"`
	Penerima_id    uint   `json:"penerima_id" from:"penerima_id"`
	Tanggal_Donasi string `json:"tanggal_donasi" form:"tanggal_donasi"`
	Qty            string `json:"qty" form:"qty"`
}
