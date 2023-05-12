package controllers

import (
	"fmt"
	"mini_project/middleware"
	"mini_project/models/payload"
	"mini_project/usecase"
	"mini_project/util"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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
			"message":          "record not found",
			"errorDescription": err.Error(),
		})
	}

	// generate jwt
	token, err := middleware.CreateToken(int(resp.ID))
	if err != nil {
		fmt.Println("LoginUserController: Error generating token")
		return err
	}

	util.CreateJWTCookies(c, token, "jwt")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success!",
		"user":    resp,
	})
}

func Authorization(c echo.Context) uint {
	cookie, _ := c.Cookie("jwt")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	id, ok := claims["user_id"].(float64)
	if !ok {
		return 0
	}
	userID := uint(id)
	return userID
}
