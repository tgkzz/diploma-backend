package course

import (
	"context"
	"server/internal/model"
	"server/internal/repository/auth"
	"server/internal/repository/course"
)

type CourseService struct {
	courseRepo course.ICourseRepo
	authRepo   auth.IAuthRepo
}

type ICourseService interface {
	GetAllCourses(ctx context.Context) ([]model.Course, error)
	GetCourse(ctx context.Context, id string) (model.Course, error)
	GetCoursesLimited(ctx context.Context) ([]model.GetCourseLimitedResponse, error)
	GetCourseLimited(ctx context.Context, id string) (model.GetCourseLimitedResponse, error)
	BuyCourse(ctx context.Context, courseId, email string) error
	CheckCourseAccess(userId int, courseId string) (bool, error)
}

func NewCourseService(courseRepo course.ICourseRepo, authRepo auth.IAuthRepo) *CourseService {
	return &CourseService{courseRepo: courseRepo, authRepo: authRepo}
}

func (c CourseService) GetAllCourses(ctx context.Context) ([]model.Course, error) {
	return c.courseRepo.GetCourses(ctx)
}

func (c CourseService) GetCoursesLimited(ctx context.Context) ([]model.GetCourseLimitedResponse, error) {
	courses, err := c.courseRepo.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	var result []model.GetCourseLimitedResponse
	for _, course := range courses {
		result = append(result, course.ToRespModel())
	}

	return result, nil
}

func (c CourseService) GetCourse(ctx context.Context, id string) (model.Course, error) {
	course, err := c.courseRepo.GetCourse(ctx, id)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (c CourseService) GetCourseLimited(ctx context.Context, id string) (model.GetCourseLimitedResponse, error) {
	course, err := c.courseRepo.GetCourse(ctx, id)
	if err != nil {
		return model.GetCourseLimitedResponse{}, err
	}

	return course.ToRespModel(), err
}

func (c CourseService) BuyCourse(ctx context.Context, courseId, email string) error {
	user, err := c.authRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	course, err := c.courseRepo.GetCourse(ctx, courseId)
	if err != nil {
		return err
	}

	if err := c.courseRepo.CreateCourseTransaction(model.Transaction{
		CourseId: courseId,
		UserId:   user.Id,
		Cost:     course.Cost,
	}); err != nil {
		return err
	}

	return nil
}

func (c CourseService) CheckCourseAccess(userId int, courseId string) (bool, error) {
	return c.courseRepo.CheckCourseAccess(userId, courseId)
}
