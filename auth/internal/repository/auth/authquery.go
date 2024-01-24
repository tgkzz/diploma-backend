package auth

import (
	"auth/internal/models"
)

// CREATE
func (a AuthRepo) CreateUser(user models.User) error {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"

	if _, err := a.DB.Exec(query, user.Username, user.Password, user.Email); err != nil {
		return err
	}

	return nil
}

// READ
func (a AuthRepo) GetUserByUsername(username string) (models.User, error) {
	query := "SELECT username, password, email from users WHERE username = $1"

	var user models.User

	err := a.DB.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Email)

	return user, err
}

//UPDATE

// DELETE
func (a AuthRepo) DeleteUserByUsername(username string) error {
	query := "DELETE FROM users WHERE username = $1"

	if _, err := a.DB.Exec(query, username); err != nil {
		return err
	}

	return nil
}
