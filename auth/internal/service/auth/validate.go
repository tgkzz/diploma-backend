package auth

import (
	"auth/internal/models"
	"auth/internal/pkg"
)

func validateUserData(user models.User) error {
	if !pkg.IsEmailValid(user.Email) {
		return models.ErrInvalidEmail
	}

	if !pkg.IsNameValid(user.FirstName, user.LastName) {
		return models.ErrInvalidName
	}

	if !pkg.IsPasswordStrong(user.Password) {
		return models.ErrInvalidPassword
	}

	return nil
}
