package models

import (
	"gorm.io/gorm"
)

type Riwayat struct {
	gorm.Model
	ID                  uint   `json:"id" from:"id"`
	Pendonor_id         uint   `json:"pendonor_id" from:"pendonor_id"`
	Tanggal_Pemeriksaan string `json:"tanggal_pemeriksaan" form:"tanggal_pemeriksaan"`
	Hasil_Pemeriksaan   string `json:"hasil_pemeriksaan" form:"hasil_pemeriksaan"`
	Golongan_Darah      string `json:"golongan_darah" form:"golongan_darah"`
}
