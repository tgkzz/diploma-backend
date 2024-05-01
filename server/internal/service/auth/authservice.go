package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"server/internal/model"
	"server/internal/pkg"
	"server/internal/repository/auth"
	"time"
)

type AuthService struct {
	repo      auth.IAuthRepo
	secretKey string
}

type IAuthService interface {
	CreateNewUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	DeleteUserByEmail(email string) error
	CheckUserCreds(creds model.User) (model.User, error)
	JwtAuthorization(user model.User) (string, error)
	//Login(user models.User) (string, error)
}

func NewAuthService(repo auth.IAuthRepo, secretKey string) *AuthService {
	return &AuthService{repo: repo, secretKey: secretKey}
}

func validateUserData(user model.User) error {
	if !pkg.IsValid(user) {
		return model.ErrEmptyness
	}

	if !pkg.IsEmailValid(user.Email) {
		return model.ErrInvalidEmail
	}

	if !pkg.IsNameValid(user.FirstName, user.LastName) {
		return model.ErrInvalidName
	}

	if !pkg.IsPasswordStrong(user.Password) {
		return model.ErrInvalidPassword
	}

	return nil
}

func (a AuthService) CreateNewUser(user model.User) (err error) {
	if err = validateUserData(user); err != nil {
		return err
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return a.repo.CreateUser(user)
}

func (a AuthService) GetUserByEmail(email string) (model.User, error) {
	return a.repo.GetUserByEmail(email)
}

func (a AuthService) DeleteUserByEmail(email string) error {
	return a.repo.DeleteUserByEmail(email)
}

func (a AuthService) CheckUserCreds(creds model.User) (model.User, error) {
	user, err := a.repo.GetUserByEmail(creds.Email)
	if err != nil {
		return model.User{}, err
	}

	if !pkg.CheckPasswordHash(creds.Password, user.Password) {
		return model.User{}, model.ErrIncorrectEmailOrPassword
	}

	return user, nil
}

func (a AuthService) JwtAuthorization(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
			ID:        pkg.RandStringBytesMaskImpr(40),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
