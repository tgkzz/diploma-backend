package course

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

const (
	DbName           = "hackaton"
	CourseCollection = "course"
)

type CourseRepo struct {
	client *mongo.Client
	db     *mongo.Database
}

type ICourseRepo interface {
	GetCourse(ctx context.Context, id string) (model.Course, error)
	GetCourses(ctx context.Context) ([]model.Course, error)
}

func NewCourseService(client *mongo.Client) *CourseRepo {
	return &CourseRepo{
		client: client,
		db:     client.Database(DbName),
	}
}

func (c CourseRepo) GetCourse(ctx context.Context, id string) (model.Course, error) {
	coll := c.db.Collection(CourseCollection)

	filter := bson.M{"_id": id}

	var course model.Course
	if err := coll.FindOne(ctx, filter).Decode(&course); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Course{}, model.ErrNotFound
		}
		return model.Course{}, err
	}

	return course, nil
}

func (c CourseRepo) GetCourses(ctx context.Context) ([]model.Course, error) {
	coll := c.db.Collection("courses")

	var courses []model.Course
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &courses); err != nil {
		return nil, err
	}

	return courses, nil
}
