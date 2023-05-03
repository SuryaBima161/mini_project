package controllers

import (
	"mini_project/lib/database"
	"mini_project/middleware"
	"mini_project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	authUser, err := database.LoginUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to login",
			"error":   err.Error(),
			"user":    authUser,
		})
	}

	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid email or password",
			"user":    nil,
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create token",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success!",
		"token":   token,
	})

}
