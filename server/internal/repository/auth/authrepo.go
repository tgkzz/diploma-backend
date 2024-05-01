package auth

import (
	"database/sql"
	"server/internal/model"
)

type AuthRepo struct {
	DB *sql.DB
}

type IAuthRepo interface {
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	DeleteUserByEmail(email string) error
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

// CREATE
func (a AuthRepo) CreateUser(user model.User) error {
	query := "INSERT INTO users (email, password, fname, lname) VALUES ($1, $2, $3, $4)"

	if _, err := a.DB.Exec(
		query,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
	); err != nil {
		return err
	}

	return nil
}

// READ
func (a AuthRepo) GetUserByEmail(email string) (model.User, error) {
	query := "SELECT id, email, password, fname, lname from users WHERE email = $1"

	var user model.User

	err := a.DB.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName)

	return user, err
}

//UPDATE

// DELETE
func (a AuthRepo) DeleteUserByEmail(email string) error {
	query := "DELETE FROM users WHERE email = $1"

	if _, err := a.DB.Exec(query, email); err != nil {
		return err
	}

	return nil
}
