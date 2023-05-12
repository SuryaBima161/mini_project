package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Id       uuid.UUID `gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Token    string `gorm:"-"`
}
