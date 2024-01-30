package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host      string `env:"HOST"`
	Port      string `env:"PORT"`
	SecretKey string `env:"SECRET_KEY"`
	DB        DB
}

type DB struct {
	DriverName     string `env:"DRIVERNAME"`
	DataSourceName string `env:"DATASOURCENAME"`
}

func LoadConfig(path string) (Config, error) {

	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host:      os.Getenv("HOST"),
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DB: DB{
			DriverName:     os.Getenv("DRIVERNAME"),
			DataSourceName: os.Getenv("DATASOURCENAME"),
		},
	}

	return cfg, nil
}
