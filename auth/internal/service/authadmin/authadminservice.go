package authadmin

import (
	"auth/internal/models"
	"auth/internal/repository/authadmin"
)

type AuthAdminService struct {
	repo authadmin.IAdminRepo
}

type IAuthAdminService interface {
	CreateNewAdmin(admin models.Admin) error
	DeleteAdmin(username string) error
	GetAdminByUsername(username string) (models.Admin, error)
	CheckAdminCreds(admin models.Admin) (models.Admin, error)
	JwtAdminAuthorization(admin models.Admin) (string, error)
}

func NewAuthService(repo authadmin.IAdminRepo) *AuthAdminService {
	return &AuthAdminService{repo: repo}
}
