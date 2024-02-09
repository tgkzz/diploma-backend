package course

import "content/internal/model"

func (c CourseService) CreateNewCourse(course model.Course) error {
	return c.repo.CreateCourse(course)
}

func (c CourseService) GetCourseByName(courseName string) (model.Course, error) {
	return c.repo.GetCourseByName(courseName)
}

func (c CourseService) GetCourseById(id int) (model.Course, error) {
	return c.repo.GetCourseById(id)
}
