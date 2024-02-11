package repository

import (
	"database/sql"
	"fakepayment/internal/repository/pay"
)

type Repository struct {
	pay.IPayRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IPayRepo: pay.NewPayRepo(db),
	}
}
