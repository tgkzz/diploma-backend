package authadmin

import (
	"auth/internal/models"
	"database/sql"
)

type AdminRepo struct {
	DB *sql.DB
}

type IAdminRepo interface {
	CreateAdmin(admin models.Admin) error
	GetAdmin(username string) (models.Admin, error)
	DeleteAdmin(username string) error
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}
