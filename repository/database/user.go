package database

import (
	"mini_project/config"
	"mini_project/models"
	"mini_project/util"
)

func LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(id uint) (user *models.User, err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	if err = config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
