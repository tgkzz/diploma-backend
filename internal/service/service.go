package service

import (
	"github.com/redis/go-redis/v9"
	"server/internal/config"
	"server/internal/repository"
	"server/internal/service/auth"
	"server/internal/service/authadmin"
	"server/internal/service/authexpert"
	"server/internal/service/course"
	"server/internal/service/meeting"
)

type Service struct {
	Auth       auth.IAuthService
	AdminAuth  authadmin.IAuthAdminService
	ExpertAuth authexpert.IExpertService
	Course     course.ICourseService
	Meeting    meeting.IMeetingService
}

func NewService(repo repository.Repository, secret string, mailCfg config.Mailer, client *redis.Client) *Service {
	return &Service{
		Auth:       auth.NewAuthService(repo, secret, mailCfg, client),
		AdminAuth:  authadmin.NewAuthService(repo, secret),
		ExpertAuth: authexpert.NewExpertService(repo, secret),
		Course:     course.NewCourseService(repo, repo),
		Meeting:    meeting.NewMeetingService(repo, repo, repo),
	}
}
