package repository

import (
	"database/sql"
	"os"
	"server/internal/config"
	"strings"

	_ "github.com/lib/pq"
)

var (
	sqlPath = "./sql/init.sql"
)

func NewDB(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.DriverName, config.DB.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return db, err
}

func createTable(db *sql.DB) error {
	fileSql, err := os.ReadFile(sqlPath)
	if err != nil {
		return err
	}

	requests := strings.Split(string(fileSql), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			return err
		}
	}
	return err
}
