package authadmin

import (
	"auth/internal/models"
	"auth/internal/pkg"
)

func ValidateAdminData(admin models.Admin) error {
	if !pkg.IsValid(admin) {
		return models.ErrEmptyness
	}

	if pkg.IsPasswordStrong(admin.Password) {
		return models.ErrInvalidPassword
	}

	return nil
}
