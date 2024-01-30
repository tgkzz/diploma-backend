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

func NewService(repo repository.Repository, secret string) *Service {
	return &Service{
		Auth:       auth.NewAuthService(repo, secret),
		AdminAuth:  authadmin.NewAuthService(repo, secret),
		ExpertAuth: authexpert.NewExpertService(repo, secret),
	}
}
