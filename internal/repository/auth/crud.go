package auth

import "diploma/internal/model/user"

func (a *AuthRepo) CreateUser(user user.User) error {
	query := "INSERT INTO users (login, password, fname, lname, email, role) Values ($1, $2, $3, $4, $5, $6)"

	if _, err := a.DB.Exec(query, user.Login, user.Password, user.FirstName, user.LastName, user.Email, user.Role); err != nil {
		return err
	}

	return nil
}

func (a *AuthRepo) GetUserByLogin(login string) (user.User, error) {
	res := user.User{}

	query := "SELECT * FROM users WHERE login = $1"

	err := a.DB.QueryRow(query, login).Scan(&res.Id, &res.Login, &res.Password, &res.Email, &res.FirstName, &res.LastName, &res.Role)

	return res, err
}
