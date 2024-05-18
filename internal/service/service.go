package service

import (
	"github.com/redis/go-redis/v9"
	"server/internal/repository"
	"server/internal/service/auth"
	"server/internal/service/authadmin"
	"server/internal/service/authexpert"
	"server/internal/service/course"
)

type Service struct {
	Auth       auth.IAuthService
	AdminAuth  authadmin.IAuthAdminService
	ExpertAuth authexpert.IExpertService
	Course     course.ICourseService
}

func NewService(repo repository.Repository, secret string, mailsenderSecret string, client *redis.Client) *Service {
	return &Service{
		Auth:       auth.NewAuthService(repo, secret, mailsenderSecret, client),
		AdminAuth:  authadmin.NewAuthService(repo, secret),
		ExpertAuth: authexpert.NewExpertService(repo, secret),
		Course:     course.NewCourseService(repo, repo),
	}
}
