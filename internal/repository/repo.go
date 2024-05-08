package repository

import (
	"database/sql"
	"server/internal/repository/auth"
	"server/internal/repository/authadmin"
	"server/internal/repository/authexpert"
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
