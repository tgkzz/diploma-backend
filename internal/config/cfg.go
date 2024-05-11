package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host          string `env:"HOST"`
	Port          string `env:"PORT"`
	SecretKey     string `env:"JWT_SECRET_KEY"`
	MailSenderKey string `env:"MAILERSEND_KEY"`
	DB            DB
	Redis         Redis
}

type DB struct {
	DriverName     string `env:"DRIVER_NAME"`
	DataSourceName string `env:"DATA_SOURCE_NAME"`
}

type Redis struct {
	Addr     string `env:"REDIS_ADDR"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (Config, error) {

	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host:          os.Getenv("HOST"),
		Port:          os.Getenv("PORT"),
		SecretKey:     os.Getenv("JWT_SECRET_KEY"),
		MailSenderKey: os.Getenv("MAILERSEND_KEY"),
		DB: DB{
			DriverName:     os.Getenv("DRIVER_NAME"),
			DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		},
		Redis: Redis{
			Addr:     os.Getenv("REDIS_ADDR"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	return cfg, nil
}
