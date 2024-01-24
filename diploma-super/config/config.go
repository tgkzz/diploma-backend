package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	return cfg, nil

}
