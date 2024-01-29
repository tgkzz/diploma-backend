package authadmin

import (
	"auth/internal/models"
	"auth/internal/pkg"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (a AuthAdminService) CreateNewAdmin(admin models.Admin) error {
	if err := ValidateAdminData(admin); err != nil {
		return err
	}

	var err error
	admin.Password, err = pkg.HashPassword(admin.Password)
	if err != nil {
		return err
	}

	if err := a.repo.CreateAdmin(admin); err != nil {
		return err
	}

	return nil
}

func (a AuthAdminService) DeleteAdmin(username string) error {
	if err := a.repo.DeleteAdmin(username); err != nil {
		return err
	}

	return nil
}

func (a AuthAdminService) GetAdminByUsername(username string) (models.Admin, error) {
	result, err := a.repo.GetAdmin(username)
	if err != nil {
		return models.Admin{}, err
	}

	return result, nil
}

func (a AuthAdminService) CheckAdminCreds(admin models.Admin) (models.Admin, error) {
	user, err := a.repo.GetAdmin(admin.Username)
	if err != nil {
		return models.Admin{}, err
	}

	if !pkg.CheckPasswordHash(admin.Password, user.Password) {
		return models.Admin{}, models.ErrIncorrectUsernameOrPassword
	}

	return user, nil
}

func (a AuthAdminService) JwtAdminAuthorization(admin models.Admin) (string, error) {
	claims := models.JwtAdminClaims{
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			ID:        pkg.RandStringBytesMaskImpr(40),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("super-puper-secret-key"))
	if err != nil {
		return "", err
	}

	return t, nil
}
