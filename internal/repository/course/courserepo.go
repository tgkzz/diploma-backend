package course

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

const (
	DbName           = "hackaton"
	CourseCollection = "courses"
)

type CourseRepo struct {
	pgDb   *sql.DB
	client *mongo.Client
	db     *mongo.Database
}

type ICourseRepo interface {
	GetCourse(ctx context.Context, id string) (model.Course, error)
	GetCourses(ctx context.Context) ([]model.Course, error)
	CreateCourseTransaction(tr model.Transaction) error
	CheckCourseAccess(userID int, courseID string) (bool, error)
	GetUserCourses(userId int) ([]string, error)
	GetCoursesContent(ctx context.Context, courseIds []string) ([]model.Course, error)
}

func NewCourseService(client *mongo.Client, pgdb *sql.DB) *CourseRepo {
	return &CourseRepo{
		client: client,
		db:     client.Database(DbName),
		pgDb:   pgdb,
	}
}

func (c CourseRepo) GetCourse(ctx context.Context, id string) (model.Course, error) {
	coll := c.db.Collection(CourseCollection)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Course{}, err
	}

	filter := bson.M{"_id": oid}

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

func (c CourseRepo) CreateCourseTransaction(tr model.Transaction) error {
	query := `INSERT INTO course_transactions (user_id, course_id, cost) VALUES ($1, $2, $3)`
	if _, err := c.pgDb.Exec(query, tr.UserId, tr.CourseId, tr.Cost); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return model.ErrCourseAlreadyPurchased
		}
		return err
	}

	return nil
}

func (c CourseRepo) GetCoursesContent(ctx context.Context, courseIds []string) ([]model.Course, error) {
	var courses []model.Course
	var objIDs []primitive.ObjectID
	coll := c.db.Collection("courses")

	for _, id := range courseIds {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIDs = append(objIDs, objID)
	}

	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c CourseRepo) CheckCourseAccess(userID int, courseID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM course_transactions WHERE user_id=$1 AND course_id=$2)`
	err := c.pgDb.QueryRow(query, userID, courseID).Scan(&exists)
	return exists, err
}

func (c CourseRepo) GetUserCourses(userId int) ([]string, error) {
	query := `SELECT course_id FROM course_transactions WHERE user_id = $1`

	rows, err := c.pgDb.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []string
	for rows.Next() {
		var courseID string
		if err := rows.Scan(&courseID); err != nil {
			return nil, err
		}
		courses = append(courses, courseID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
