package pkg

import (
	"regexp"
)

func IsPasswordStrong(password string) bool {
	var (
		hasMinLen      = regexp.MustCompile(`.{8,}`)
		hasNumber      = regexp.MustCompile(`[0-9]+`)
		hasUpper       = regexp.MustCompile(`[A-Z]+`)
		hasLower       = regexp.MustCompile(`[a-z]+`)
		hasSpecialChar = regexp.MustCompile(`[!@#\$%\^&\*\(\)_]+`)
	)
	if !hasMinLen.MatchString(password) {
		return false
	}
	if !hasNumber.MatchString(password) {
		return false
	}
	if !hasUpper.MatchString(password) {
		return false
	}
	if !hasLower.MatchString(password) {
		return false
	}
	if !hasSpecialChar.MatchString(password) {
		return false
	}
	return true
}
