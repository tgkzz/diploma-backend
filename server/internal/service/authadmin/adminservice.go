package authadmin

import (
	"github.com/golang-jwt/jwt/v5"
	"server/internal/model"
	"server/internal/pkg"
	"server/internal/repository/authadmin"
	"time"
)

type AuthAdminService struct {
	repo      authadmin.IAdminRepo
	secretKey string
}

type IAuthAdminService interface {
	CreateNewAdmin(admin model.Admin) error
	DeleteAdmin(username string) error
	GetAdminByUsername(username string) (model.Admin, error)
	CheckAdminCreds(admin model.Admin) (model.Admin, error)
	JwtAdminAuthorization(admin model.Admin) (string, error)
}

func NewAuthService(repo authadmin.IAdminRepo, secret string) *AuthAdminService {
	return &AuthAdminService{repo: repo, secretKey: secret}
}

func ValidateAdminData(admin model.Admin) error {
	if !pkg.IsValid(admin) {
		return model.ErrEmptyness
	}

	if pkg.IsPasswordStrong(admin.Password) {
		return model.ErrInvalidPassword
	}

	return nil
}

func (a AuthAdminService) CreateNewAdmin(admin model.Admin) error {
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

func (a AuthAdminService) GetAdminByUsername(username string) (model.Admin, error) {
	result, err := a.repo.GetAdmin(username)
	if err != nil {
		return model.Admin{}, err
	}

	return result, nil
}

func (a AuthAdminService) CheckAdminCreds(admin model.Admin) (model.Admin, error) {
	user, err := a.repo.GetAdmin(admin.Username)
	if err != nil {
		return model.Admin{}, err
	}

	if !pkg.CheckPasswordHash(admin.Password, user.Password) {
		return model.Admin{}, model.ErrIncorrectUsernameOrPassword
	}

	return user, nil
}

func (a AuthAdminService) JwtAdminAuthorization(admin model.Admin) (string, error) {
	claims := model.JwtAdminClaims{
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			ID:        pkg.RandStringBytesMaskImpr(40),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
