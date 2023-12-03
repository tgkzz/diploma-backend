package service

import (
	"diploma/internal/repository"
	"diploma/internal/service/auth"
	"diploma/internal/service/payment"
)

type Service struct {
	auth.AutherService
	payment.PaymenterService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AutherService:    auth.NewAuthService(repo),
		PaymenterService: payment.NewPaymentService(),
	}
}
