package service

import (
	"diploma/internal/repository"
	"diploma/internal/service/auth"
)

type Service struct {
	auth.AutherService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AutherService: auth.NewAuthService(repo),
	}
}
