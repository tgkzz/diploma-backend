package authadmin

import "auth/internal/models"

// CREATE
func (a AdminRepo) CreateAdmin(admin models.Admin) error {
	query := `INSERT INTO admin (username, password) VALUES ($1, $2)`

	_, err := a.DB.Exec(query, admin.Username, admin.Password)

	return err
}

// READ
func (a AdminRepo) GetAdmin(username string) (models.Admin, error) {
	query := `SELECT id, username, password FROM admin WHERE username = $1`

	var admin models.Admin

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
