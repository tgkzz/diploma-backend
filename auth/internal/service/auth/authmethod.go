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

func (a AuthService) GetUserByEmail(email string) (models.User, error) {
	return a.repo.GetUserByEmail(email)
}

func (a AuthService) DeleteUserByEmail(email string) error {
	return a.repo.DeleteUserByEmail(email)
}

func (a AuthService) CheckUserCreds(creds models.User) (models.User, error) {
	user, err := a.repo.GetUserByEmail(creds.Email)
	if err != nil {
		return models.User{}, err
	}

	if !checkPasswordHash(creds.Password, user.Password) {
		return models.User{}, models.ErrIncorrectEmailOrPassword
	}

	return user, nil
}

func (a AuthService) JwtAuthorization(user models.User) (string, error) {
	claims := &models.JwtCustomClaims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			ID:        randStringBytesMaskImpr(40),
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
