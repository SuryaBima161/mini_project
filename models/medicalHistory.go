package models

import (
	"gorm.io/gorm"
)

type MedicalHistory struct {
	gorm.Model
	DonaturID          uint   `json:"pendonor_id" from:"pendonor_id" gorm:"foreignkey:DonaturID"`
	TanggalPemeriksaan string `json:"tanggal_pemeriksaan" form:"tanggal_pemeriksaan"`
	HasilPemeriksaan   string `json:"hasil_pemeriksaan" form:"hasil_pemeriksaan"`
	GolonganDarah      string `json:"golongan_darah" form:"golongan_darah"`
}
