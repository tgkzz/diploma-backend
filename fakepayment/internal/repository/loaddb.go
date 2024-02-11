package repository

import (
	"database/sql"
	"fakepayment/config"

	_ "github.com/lib/pq"
)

func NewDB(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.DriverName, config.DB.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
