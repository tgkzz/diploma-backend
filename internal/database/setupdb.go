package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	connStr string = "postgresql://postgres:1234@localhost:5432/diploma?sslmode=disable"
)

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
