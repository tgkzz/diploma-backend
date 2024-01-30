package service

import (
	"auth/internal/repository"
	"auth/internal/service/auth"
	"auth/internal/service/authadmin"
	"auth/internal/service/authexpert"
)

type Service struct {
	Auth       auth.IAuthService
	AdminAuth  authadmin.IAuthAdminService
	ExpertAuth authexpert.IExpertService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Auth:       auth.NewAuthService(repo),
		AdminAuth:  authadmin.NewAuthService(repo),
		ExpertAuth: authexpert.NewExpertService(repo),
	}
}
