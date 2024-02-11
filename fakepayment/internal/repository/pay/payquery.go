package pay

import "fakepayment/internal/model"

func (p PayRepo) CreateNewTransaction(tr model.Transaction) error {
	query := `INSERT INTO course_transaction (user_id, course_id, cost) VALUES ($1, $2, $3)`
	_, err := p.DB.Exec(query, tr.UserId, tr.CourseId, tr.Cost)
	if err != nil {
		return err
	}
	return nil
}

func (p PayRepo) GetTransactionById(id int) (model.Transaction, error) {
	query := `SELECT id, user_id, course_id, cost FROM course_transaction WHERE id = $1`
	var tr model.Transaction

	row := p.DB.QueryRow(query, id)
	err := row.Scan(&tr.Id, &tr.UserId, &tr.CourseId, &tr.Cost)
	if err != nil {
		return model.Transaction{}, err
	}
	return tr, nil
}
