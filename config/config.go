package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
	DB   DB
}

type DB struct {
	DriverName     string
	DataSourceName string
}

func OpenConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

	config := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DB: DB{
			DriverName:     os.Getenv("DRIVERNAME"),
			DataSourceName: os.Getenv("DATASOURCENAME"),
		},
	}

	return config, nil
}
