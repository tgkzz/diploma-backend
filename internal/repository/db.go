package repository

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func NewMongoDB(cfg config.Mongo) (*mongo.Client, error) {
	ctx := context.Background()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(cfg.Uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
