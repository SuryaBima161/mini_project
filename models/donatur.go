package models

import (
	"gorm.io/gorm"
)

type Donatur struct {
	gorm.Model
	UserID              uint   `json:"user_id" form:"user_id" gorm:"foreignkey:UserID"`
	Name                string `json:"name" form:"name"`
	JenisKelamin        string `json:"jenis_kelamin" form:"jenis_kelamin"`
	TanggalLahir        string `json:"tanggal_lahir" form:"tanggal_lahir"`
	DetailBloodDonation []DetailBloodDonation
}
