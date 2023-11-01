package database

import (
	"diploma/internal/model"
	"diploma/internal/pkg"
	"regexp"
)

func CheckUserCreds(user model.User) (string, bool) {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(user.Email) {
		return "email must be valid", true
	}

	if !pkg.IsPasswordStrong(user.Password) {
		return "password is weak", true
	}

	return "", false
}
