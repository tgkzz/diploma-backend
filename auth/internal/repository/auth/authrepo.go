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
	GetUserByUsername(username string) (models.User, error)
	DeleteUserByUsername(username string) error
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}
