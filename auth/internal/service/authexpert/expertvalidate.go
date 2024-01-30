package authexpert

import (
	"auth/internal/models"
	"auth/internal/pkg"
)

func validateExpertData(expert models.Expert) error {
	if !pkg.IsValid(expert) {
		return models.ErrEmptyness
	}

	if !pkg.IsEmailValid(expert.Email) {
		return models.ErrInvalidEmail
	}

	if !pkg.IsNameValid(expert.FirstName, expert.LastName) {
		return models.ErrInvalidName
	}

	if !pkg.IsPasswordStrong(expert.Password) {
		return models.ErrInvalidPassword
	}

	return nil
}
