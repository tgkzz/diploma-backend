package auth

import (
	"auth/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (a AuthService) CreateNewUser(user models.User) (err error) {
	if err = validateUserData(user); err != nil {
		return err
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	return a.repo.CreateUser(user)
}

func (a AuthService) GetUserByUsername(username string) (models.User, error) {
	return a.repo.GetUserByUsername(username)
}

func (a AuthService) DeleteUserByUsername(username string) error {
	return a.repo.DeleteUserByUsername(username)
}

func (a AuthService) CheckUserCreds(creds models.User) (models.User, error) {
	user, err := a.repo.GetUserByUsername(creds.Username)
	if err != nil {
		return models.User{}, err
	}

	if !checkPasswordHash(creds.Password, user.Password) {
		return models.User{}, models.ErrIncorrectUsernameOrEmail
	}

	return user, nil
}

func (a AuthService) JwtAuthorization(user models.User) (string, error) {
	claims := &models.JwtCustomClaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//TODO: implement normal key for signing
	t, err := token.SignedString([]byte("super-puper-secret-key"))
	if err != nil {
		return "", err
	}

	return t, nil
}
