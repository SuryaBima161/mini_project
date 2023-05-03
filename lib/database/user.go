package database

import (
	"mini_project/config"
	"mini_project/models"
)

// baru
func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return err, nil
	}
	return user, nil
}

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(id uint) (user models.User, err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {

	if e := config.DB.Create(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}

func UpdateUser(user models.User, id uint) error {
	if e := config.DB.Updates(&user).Where("id = ?", id).Error; e != nil {
		return e
	}
	return nil
}

func DeleteUser(id uint) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return id, err
	}
	return id, nil
}
