package auth

import (
	"database/sql"
	"diploma/internal/model/user"
)

type AuthRepo struct {
	DB *sql.DB
}

type AutherRepo interface {
	//create
	CreateUser(user user.User) error
	//read
	GetUserByLogin(login string) (user.User, error)
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		DB: db,
	}
}
