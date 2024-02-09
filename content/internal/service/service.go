package service

import (
	"content/internal/repository"
	"content/internal/service/course"
)

type Service struct {
	Course course.ICourseService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Course: course.NewCourseService(repo),
	}
}
