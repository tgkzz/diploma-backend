package service

import "super/internal/service/auth"

type Service struct {
	Auth auth.IAuthService
}

func NewService(AuthURL string) *Service {
	return &Service{
		Auth: auth.NewAuthService(AuthURL),
	}
}
