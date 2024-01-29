package pkg

import (
	"regexp"
	"strings"
)

func IsNameValid(firstName, lastName string) bool {
	return len(strings.TrimSpace(firstName)) != 0 && len(strings.TrimSpace(lastName)) != 0
}

func IsEmailValid(email string) bool {
	asd := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)
	return asd
}

func IsPasswordStrong(password string) bool {
	var (
		hasMinLen = len(password) >= 8
		hasUpper  = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower  = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber = regexp.MustCompile(`[0-9]`).MatchString(password)
	)
	return hasMinLen && hasUpper && hasLower && hasNumber
}
