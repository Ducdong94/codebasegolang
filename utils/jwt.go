package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
