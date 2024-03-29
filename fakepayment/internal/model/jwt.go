package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}
