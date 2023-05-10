package models

import (
	"gorm.io/gorm"
)

type Pendonor struct {
	gorm.Model
	ID            uint   `json:"id" from:"id"`
	User_id       uint   `json:"user_id" form:"user_id"`
	Name          string `json:"name" form:"name"`
	Jenis_Kelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	Tanggal_lahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
}
