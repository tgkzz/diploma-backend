package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host      string `env:"HOST"`
	Port      string `env:"PORT"`
	SecretKey string `env:"JWT_SECRET_KEY"`
	DB        DB
}

type DB struct {
	DriverName     string `env:"DRIVER_NAME"`
	DataSourceName string `env:"DATA_SOURCE_NAME"`
}

func LoadConfig(path string) (Config, error) {

	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host:      os.Getenv("HOST"),
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		DB: DB{
			DriverName:     os.Getenv("DRIVER_NAME"),
			DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		},
	}

	return cfg, nil
}
