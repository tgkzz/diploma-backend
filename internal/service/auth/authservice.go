package auth

import (
	"diploma/internal/model/user"
	"diploma/internal/repository/auth"
)

type AuthService struct {
	repo auth.AutherRepo
}

type AutherService interface {
	CreateUser(user user.User) error
	CheckUserCreds(creds user.User) error
	CreateToken(login string) (string, error)
}

func NewAuthService(repo auth.AutherRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
