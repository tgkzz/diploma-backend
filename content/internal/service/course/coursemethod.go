package course

import (
	"content/internal/model"
	"content/internal/pkg"
)

func (c CourseService) CreateNewCourse(course model.Course) error {
	return c.repo.CreateCourse(course)
}

func (c CourseService) GetCourseByName(courseName string) (model.Course, error) {
	return c.repo.GetCourseByName(courseName)
}

func (c CourseService) GetCourseById(id string) (model.Course, error) {
	idInt, err := pkg.StrictAtoi(id)
	if err != nil {
		return model.Course{}, err
	}

	return c.repo.GetCourseById(idInt)
}

func (c CourseService) GetAllCourse() ([]model.Course, error) {
	return c.repo.GetAllCourse()
}
