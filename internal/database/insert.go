package database

import (
	"diploma/internal/model"
	"errors"
)

func InsertUser(user model.User) error {
	db, err := setupDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if user == (model.User{}) {
		return errors.New("empty struct came into input")
	}

	query := "INSERT INTO users(login, password, email, fname, lname, role) VALUES($1, $2, $3, $4, $5, $6)"

	_, err = db.Exec(query, user.Login, user.Password, user.Email, user.FirstName, user.LastName, user.Role)
	if err != nil {
		return err
	}

	return nil
}
