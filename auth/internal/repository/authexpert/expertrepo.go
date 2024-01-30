package authexpert

import (
	"auth/internal/models"
	"database/sql"
)

type ExpertRepo struct {
	DB *sql.DB
}

type IExpertRepo interface {
	CreateExpert(expert models.Expert) error
	DeleteExpertByEmail(email string) error
	GetExpertByEmail(email string) (models.Expert, error)
}

func NewExpertRepo(db *sql.DB) *ExpertRepo {
	return &ExpertRepo{
		DB: db,
	}
}
