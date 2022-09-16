package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"os"
	"time"
)

type TokenData struct {
	UserId uint
	IsAuthorized bool
}

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractTokenData(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid{
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}
