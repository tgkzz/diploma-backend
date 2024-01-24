package auth

import (
	"auth/internal/models"
	"regexp"
)

func validateUserData(user models.User) error {
	if !isEmailValid(user.Email) {
		return models.ErrInvalidEmail
	}

	if !isPasswordStrong(user.Password) {
		return models.ErrInvalidPassword
	}

	return nil
}

func isEmailValid(email string) bool {
	asd := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)
	return asd
}

func isPasswordStrong(password string) bool {
	var (
		hasMinLen = len(password) >= 8
		hasUpper  = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower  = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber = regexp.MustCompile(`[0-9]`).MatchString(password)
	)
	return hasMinLen && hasUpper && hasLower && hasNumber
}
