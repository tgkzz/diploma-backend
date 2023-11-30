package repository

import (
	"database/sql"
	"diploma/config"

	_ "github.com/lib/pq"
)

func NewDB(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.DriverName, config.DB.DataSourceName)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, err
}
