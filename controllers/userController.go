package controllers

import (
	"mini_project/models"
	"mini_project/models/payload"
	"mini_project/usecase"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// get user by id
func GetUsersIdController(c echo.Context) error {
	// id := Authorization(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetUserById(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error get user",
			"errorDescription": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

// create user
func CreateUserController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create new user",
			"errorDescription": err.Error(),
		})
	}
	respone, err := usecase.CreateUser(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create new user",
			"errorDescription": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    respone,
	})
}

func UpdateUserController(c echo.Context) error {
	// id := Authorization(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := models.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update user",
			"errorDescription": err.Error(),
			"errorMessage":     "Mohon Maaf user tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

// delete user
func DeleteUserController(c echo.Context) error {
	// id := Authorization(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": err.Error(),
			"errorMessage":     "Mohon Maaf user tidak dapat di delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
