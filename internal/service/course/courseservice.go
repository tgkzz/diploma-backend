package course

import (
	"context"
	"server/internal/model"
	"server/internal/repository/course"
)

type CourseService struct {
	courseRepo course.ICourseRepo
}

type ICourseService interface {
	GetAllCourses(ctx context.Context) ([]model.Course, error)
}

func NewCourseService(courseRepo course.ICourseRepo) *CourseService {
	return &CourseService{courseRepo: courseRepo}
}

func (c CourseService) GetAllCourses(ctx context.Context) ([]model.Course, error) {
	return c.courseRepo.GetCourses(ctx)
}
