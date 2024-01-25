package auth

import (
	"auth/internal/models"
)

// CREATE
func (a AuthRepo) CreateUser(user models.User) error {
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
func (a AuthRepo) GetUserByEmail(email string) (models.User, error) {
	query := "SELECT email, password, fname, lname from users WHERE email = $1"

	var user models.User

	err := a.DB.QueryRow(query, email).Scan(&user.Email, &user.Password, &user.FirstName, &user.LastName)

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
