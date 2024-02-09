package course

import (
	"content/internal/model"
	"database/sql"
)

type CourseRepo struct {
	DB *sql.DB
}

type ICourseRepo interface {
	CreateCourse(course model.Course) error
	GetCourseByName(courseName string) (model.Course, error)
	GetCourseById(id int) (model.Course, error)
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}
