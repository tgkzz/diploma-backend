package repository

import (
	"database/sql"
	"diploma/internal/repository/auth"
)

type Repository struct {
	auth.AutherRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AutherRepo: auth.NewAuthRepo(db),
	}
}
