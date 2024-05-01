package service

import (
	"server/internal/repository"
	"server/internal/service/auth"
	"server/internal/service/authadmin"
	"server/internal/service/authexpert"
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
