package auth

import (
	"auth/internal/models"
	"auth/internal/repository/auth"
)

type AuthService struct {
	repo auth.IAuthRepo
}

type IAuthService interface {
	CreateNewUser(user models.User) error
	GetUserByUsername(username string) (models.User, error)
	DeleteUserByUsername(username string) error
	CheckUserCreds(creds models.User) (models.User, error)
	JwtAuthorization(user models.User) (string, error)
	//Login(user models.User) (string, error)
}

func NewAuthService(repo auth.IAuthRepo) *AuthService {
	return &AuthService{repo: repo}
}
