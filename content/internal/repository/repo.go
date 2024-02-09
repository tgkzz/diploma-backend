package repository

import (
	"content/internal/repository/course"
	"database/sql"
)

type Repository struct {
	course.ICourseRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ICourseRepo: course.NewCourseRepo(db),
	}
}
