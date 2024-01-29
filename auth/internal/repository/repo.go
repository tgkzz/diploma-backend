package repository

import (
	"auth/internal/repository/auth"
	"auth/internal/repository/authadmin"
	"database/sql"
)

type Repository struct {
	auth.IAuthRepo
	authadmin.IAdminRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IAuthRepo:  auth.NewAuthRepo(db),
		IAdminRepo: authadmin.NewAdminRepo(db),
	}
}
