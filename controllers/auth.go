package controllers

import (
	"mini_project/models/payload"
	"mini_project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := payload.LoginUserRequest{}
	c.Bind(&user)
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error login user",
			"errorDescription": err,
		})
	}

	resp, err := usecase.LoginUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "record not found",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success!",
		"user":    resp,
	})
}
