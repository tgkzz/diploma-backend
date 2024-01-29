package service

import (
	"auth/internal/repository"
	"auth/internal/service/auth"
	"auth/internal/service/authadmin"
)

type Service struct {
	Auth      auth.IAuthService
	AdminAuth authadmin.IAuthAdminService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Auth:      auth.NewAuthService(repo),
		AdminAuth: authadmin.NewAuthService(repo),
	}
}
