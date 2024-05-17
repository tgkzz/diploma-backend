package repository

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/repository/auth"
	"server/internal/repository/authadmin"
	"server/internal/repository/authexpert"
	"server/internal/repository/course"
)

type Repository struct {
	auth.IAuthRepo
	authadmin.IAdminRepo
	authexpert.IExpertRepo
	course.ICourseRepo
}

func NewRepository(db *sql.DB, client *mongo.Client) *Repository {
	return &Repository{
		IAuthRepo:   auth.NewAuthRepo(db),
		IAdminRepo:  authadmin.NewAdminRepo(db),
		IExpertRepo: authexpert.NewExpertRepo(db),
		ICourseRepo: course.NewCourseService(client),
	}
}
