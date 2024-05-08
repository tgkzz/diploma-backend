package authexpert

import (
	"database/sql"
	"server/internal/model"
)

type ExpertRepo struct {
	DB *sql.DB
}

type IExpertRepo interface {
	CreateExpert(expert model.Expert) error
	DeleteExpertByEmail(email string) error
	GetExpertByEmail(email string) (model.Expert, error)
	GetAllExpert() ([]model.Expert, error)
}

func NewExpertRepo(db *sql.DB) *ExpertRepo {
	return &ExpertRepo{
		DB: db,
	}
}

func (e ExpertRepo) CreateExpert(expert model.Expert) error {
	query := `INSERT INTO expert (fname, lname, email, cost, password) VALUES ($1, $2, $3, $4, $5)`

	_, err := e.DB.Exec(query, expert.FirstName, expert.LastName, expert.Email, expert.Cost, expert.Password)

	return err

}

func (e ExpertRepo) GetAllExpert() ([]model.Expert, error) {
	var experts []model.Expert

	query := `SELECT id, fname, lname, email, cost, description FROM expert`
	rows, err := e.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expert model.Expert
		err := rows.Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Description)
		if err != nil {
			return nil, err
		}
		experts = append(experts, expert)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return experts, nil
}

func (e ExpertRepo) DeleteExpertByEmail(email string) error {
	query := `DELETE FROM expert WHERE email = $1`

	_, err := e.DB.Exec(query, email)

	return err
}

func (e ExpertRepo) GetExpertByEmail(email string) (model.Expert, error) {
	query := `SELECT id, fname, lname, email, cost, password FROM expert WHERE email = $1`

	var expert model.Expert

	err := e.DB.QueryRow(query, email).Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Password)

	return expert, err
}
