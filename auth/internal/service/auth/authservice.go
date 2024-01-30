package auth

import (
	"auth/internal/models"
	"auth/internal/repository/auth"
)

type AuthService struct {
	repo      auth.IAuthRepo
	secretKey string
}

type IAuthService interface {
	CreateNewUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	DeleteUserByEmail(email string) error
	CheckUserCreds(creds models.User) (models.User, error)
	JwtAuthorization(user models.User) (string, error)
	//Login(user models.User) (string, error)
}

func NewAuthService(repo auth.IAuthRepo, secretKey string) *AuthService {
	return &AuthService{repo: repo, secretKey: secretKey}
}
