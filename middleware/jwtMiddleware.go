package middleware

import (
	cons "mini_project/constant"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	// jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorize"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cons.SECRET_JWT))
}

var Cookie = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(cons.SECRET_JWT),
	TokenLookup:   "cookie:jwt",
	AuthScheme:    "user",
})
