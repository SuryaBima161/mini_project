package models

import (
	"gorm.io/gorm"
)

type DetailBloodDonation struct {
	gorm.Model
	ID                 uint   `json:"id" from:"id"`
	Pendonor_id        uint   `json:"pendonor_id" from:"pendonor_id"`
	Penerima_id        uint   `json:"penerima_id" from:"penerima_id"`
	Riwayat_id         uint   `json:"riwayat_id" form:"riwayat_id"`
	Pemberian_darah_id uint   `json:"pemberian_darah_id" form:"pemberian_darah_id"`
	Total_Qty          string `json:"qty" form:"qty"`
}
