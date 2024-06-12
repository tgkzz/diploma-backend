package auth

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"server/internal/model"
	"strings"
)

type AuthRepo struct {
	DB *sql.DB
}

type IAuthRepo interface {
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(req model.UpdateUserRequest) error
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

func (a AuthRepo) GetUserById(id int) (model.User, error) {
	query := "SELECT id, email, password, fname, lname from users WHERE id = $1"

	var user model.User

	if err := a.DB.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName); err != nil {
		return model.User{}, err
	}

	return user, nil
}

// UPDATE
func (a AuthRepo) UpdateUser(req model.UpdateUserRequest) error {
	var updates []string
	var params []interface{}

	index := 1

	if req.HasEmail() {
		updates = append(updates, fmt.Sprintf("email = $%d", index))
		params = append(params, req.Email)
		index++
	}
	if req.HasPassword() {
		updates = append(updates, fmt.Sprintf("password = $%d", index))
		params = append(params, req.Password)
		index++
	}
	if req.HasFname() {
		updates = append(updates, fmt.Sprintf("fname = $%d", index))
		params = append(params, req.Fname)
		index++
	}
	if req.HasLname() {
		updates = append(updates, fmt.Sprintf("lname = $%d", index))
		params = append(params, req.Lname)
		index++
	}

	if len(updates) == 0 {
		return nil
	}

	updatesStr := strings.Join(updates, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", updatesStr, index)
	params = append(params, req.Id)

	_, err := a.DB.Exec(query, params...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return model.ErrEmailIsAlreadyUser
		}
		return err
	}

	return nil
}

// DELETE
func (a AuthRepo) DeleteUserByEmail(email string) error {
	query := "DELETE FROM users WHERE email = $1"

	if _, err := a.DB.Exec(query, email); err != nil {
		return err
	}

	return nil
}
