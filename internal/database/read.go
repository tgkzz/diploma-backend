package database

import "diploma/internal/model"

func ReturnUser(login string) (model.User, error) {
	var id string

	Result := model.User{}

	db, err := setupDB()
	if err != nil {
		return Result, err
	}
	defer db.Close()

	query := "SELECT * FROM users WHERE login = $1"

	err = db.QueryRow(query, login).Scan(&id, &Result.Login, &Result.Password, &Result.Email, &Result.FirstName, &Result.LastName, &Result.Role)
	if err != nil {
		return Result, err
	}

	return Result, nil
}
