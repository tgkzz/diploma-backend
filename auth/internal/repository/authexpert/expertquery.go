package authexpert

import "auth/internal/models"

func (e ExpertRepo) CreateExpert(expert models.Expert) error {
	query := `INSERT INTO expert (fname, lname, email, cost, password) VALUES ($1, $2, $3, $4, $5)`

	_, err := e.DB.Exec(query, expert.FirstName, expert.LastName, expert.Email, expert.Cost, expert.Password)

	return err

}

func (e ExpertRepo) DeleteExpertByEmail(email string) error {
	query := `DELETE FROM expert WHERE email = $1`

	_, err := e.DB.Exec(query, email)

	return err
}

func (e ExpertRepo) GetExpertByEmail(email string) (models.Expert, error) {
	query := `SELECT id, fname, lname, email, cost, password FROM expert WHERE email = $1`

	var expert models.Expert

	err := e.DB.QueryRow(query, email).Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Password)

	return expert, err
}
