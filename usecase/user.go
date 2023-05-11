package usecase

import (
	"errors"
	"mini_project/config"
	"mini_project/middleware"
	"mini_project/models"
	"mini_project/models/payload"
	"mini_project/repository/database"
	"mini_project/util"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(req *payload.LoginUserRequest) (resp payload.LoginUserResponse, err error) {
	loginUser := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err = database.LoginUser(loginUser)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}

	if err := config.DB.Where(loginUser).First(loginUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return payload.LoginUserResponse{}, errors.New("Invalid email or password")
		}
		fmt.Println("Error getting user from database:", err)
		return payload.LoginUserResponse{}, err
	}

	// Pembandingan password
	if err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password)); err != nil {
		return
	}
	// generate jwts
	token, err := middleware.CreateToken(int(loginUser.ID))
	if err != nil {
		fmt.Println("GetUser: Error Generate token")
		return
	}
	resp = payload.LoginUserResponse{
		ID:    loginUser.ID,
		Email: loginUser.Email,
		Token: token,
	}
	return
}
func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	newUser := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return payload.CreateUserResponse{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Email already exists",
			"errorDescription": err,
		})
	} else if err != gorm.ErrRecordNotFound {
		return payload.CreateUserResponse{}, echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed to query database",
			"errorDescription": err,
		})
	}
	if req.Password != "" {
		hash, err := util.HashPassword(req.Password)
		if err != nil {
			return payload.CreateUserResponse{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message":          "error hash password",
				"errorDescription": err,
			})
		}
		req.Password = hash
	}
	err = database.CreateUser(newUser)
	if err != nil {
		return
	}
	resp = payload.CreateUserResponse{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	return
}

func GetUserById(id uint) (*payload.GetUserByIdResponse, error) {
	user, err := database.GetUser(id)
	if err != nil {
		return nil, err
	}
	resp := payload.GetUserByIdResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return &resp, nil
}

func UpdateUser(user *models.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}
	return
}
func DeleteUser(userId uint) (err error) {
	deleteUser := models.User{
		ID: userId,
	}
	err = database.DeleteUser(&deleteUser)
	if err != nil {
		fmt.Println("DeleteUser: error deleting user, err:", err)
		return
	}
	return err
}
