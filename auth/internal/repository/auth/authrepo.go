package auth

import (
	"auth/internal/models"
	"database/sql"
)

type AuthRepo struct {
	DB *sql.DB
}

type IAuthRepo interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	DeleteUserByEmail(email string) error
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}
