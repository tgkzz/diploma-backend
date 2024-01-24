package service

import (
	"auth/internal/repository"
	"auth/internal/service/auth"
)

type Service struct {
	Auth auth.IAuthService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Auth: auth.NewAuthService(repo),
	}
}
