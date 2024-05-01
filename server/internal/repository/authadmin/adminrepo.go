package authadmin

import (
	"database/sql"
	"server/internal/model"
)

type AdminRepo struct {
	DB *sql.DB
}

type IAdminRepo interface {
	CreateAdmin(admin model.Admin) error
	GetAdmin(username string) (model.Admin, error)
	DeleteAdmin(username string) error
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

// CREATE
func (a AdminRepo) CreateAdmin(admin model.Admin) error {
	query := `INSERT INTO admin (username, password) VALUES ($1, $2)`

	_, err := a.DB.Exec(query, admin.Username, admin.Password)

	return err
}

// READ
func (a AdminRepo) GetAdmin(username string) (model.Admin, error) {
	query := `SELECT id, username, password FROM admin WHERE username = $1`

	var admin model.Admin

	err := a.DB.QueryRow(query, username).Scan(&admin.Id, &admin.Username, &admin.Password)

	return admin, err
}

// UPDATE

// DELETE
func (a AdminRepo) DeleteAdmin(username string) error {
	query := `DELETE FROM admin WHERE username = $1`

	_, err := a.DB.Exec(query, username)

	return err
}
