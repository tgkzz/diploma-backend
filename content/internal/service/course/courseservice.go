package course

import (
	"content/internal/model"
	"content/internal/repository/course"
)

type CourseService struct {
	repo course.ICourseRepo
}

type ICourseService interface {
	CreateNewCourse(course model.Course) error
	GetCourseByName(courseName string) (model.Course, error)
	GetCourseById(id string) (model.Course, error)
	GetAllCourse() ([]model.Course, error)
}

func NewCourseService(repo course.ICourseRepo) *CourseService {
	return &CourseService{
		repo: repo,
	}
}
