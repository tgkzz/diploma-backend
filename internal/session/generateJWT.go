package session

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("secret-key")

func CreateToken() (string, error) {
	username := "Daniyar"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
