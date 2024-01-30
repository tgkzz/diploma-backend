package repository

import (
	"auth/internal/repository/auth"
	"auth/internal/repository/authadmin"
	"auth/internal/repository/authexpert"
	"database/sql"
)

type Repository struct {
	auth.IAuthRepo
	authadmin.IAdminRepo
	authexpert.IExpertRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IAuthRepo:   auth.NewAuthRepo(db),
		IAdminRepo:  authadmin.NewAdminRepo(db),
		IExpertRepo: authexpert.NewExpertRepo(db),
	}
}
