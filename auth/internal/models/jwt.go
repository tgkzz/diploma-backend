package models

import jwt "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
