package pay

import (
	"database/sql"
	"fakepayment/internal/model"
)

type PayRepo struct {
	DB *sql.DB
}

type IPayRepo interface {
	CreateNewTransaction(tr model.Transaction) error
	GetTransactionById(id int) (model.Transaction, error)
}

func NewPayRepo(db *sql.DB) *PayRepo {
	return &PayRepo{DB: db}
}
