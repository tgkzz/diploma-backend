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
	GetExpertById(id int) (model.Expert, error)
}

func NewExpertRepo(db *sql.DB) *ExpertRepo {
	return &ExpertRepo{
		DB: db,
	}
}

func (e ExpertRepo) CreateExpert(expert model.Expert) error {
	query := `INSERT INTO experts (fname, lname, email, cost, password, description, imageLink) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := e.DB.Exec(query, expert.FirstName, expert.LastName, expert.Email, expert.Cost, expert.Password, expert.Description, expert.ImageLink)

	return err

}

func (e ExpertRepo) GetAllExpert() ([]model.Expert, error) {
	var experts []model.Expert

	query := `SELECT id, fname, lname, email, cost, description, imageLink FROM experts`
	rows, err := e.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expert model.Expert
		err := rows.Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Description, &expert.ImageLink)
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
	query := `DELETE FROM experts WHERE email = $1`

	_, err := e.DB.Exec(query, email)

	return err
}

func (e ExpertRepo) GetExpertByEmail(email string) (model.Expert, error) {
	query := `SELECT id, fname, lname, email, cost, password, description, imageLink FROM experts WHERE email = $1`

	var expert model.Expert

	err := e.DB.QueryRow(query, email).Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Password, &expert.Description, &expert.ImageLink)

	return expert, err
}

func (e ExpertRepo) GetExpertById(id int) (model.Expert, error) {
	query := `SELECT id, fname, lname, email, cost, password, description, imageLink FROM experts WHERE id = $1`

	var expert model.Expert

	err := e.DB.QueryRow(query, id).Scan(&expert.Id, &expert.FirstName, &expert.LastName, &expert.Email, &expert.Cost, &expert.Password, &expert.Description, &expert.ImageLink)

	return expert, err
}
