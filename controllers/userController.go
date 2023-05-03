package controllers

import (
	"mini_project/config"
	"mini_project/lib/database"
	"mini_project/models"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// get user by id
func GetUsersIdController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := database.GetUser(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get  user",
		"users":   user,
	})
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	match := regex.MatchString(strings.TrimSpace(email))
	return match
}

// create user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// validasi email dupe
	if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "Email already exists",
			"errorDescription": err,
		})
	}

	// validasi format email
	if !isValidEmail(user.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid email format",
		})
	}

	if createdUser, err := database.CreateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create new user",
			"errorDescription": err,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    createdUser,
		})
	}
}

// update user
func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	updateUser := models.User{}
	c.Bind(&updateUser)
	if e := database.UpdateUser(updateUser, uint(userId)); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed update user",
			"errorDescription": e,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    updateUser,
	})
}

// delete user
func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	deleteUser, e := database.DeleteUser(uint(userId))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete user",
			"errorDescription": e,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    deleteUser,
	})
}
