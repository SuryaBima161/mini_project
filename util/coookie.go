package util

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateJWTCookies(c echo.Context, token string, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)
}
