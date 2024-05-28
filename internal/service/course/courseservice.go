package course

import (
	"context"
	"errors"
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
	GetUserCourses(ctx context.Context, email string) ([]model.Course, error)
	GetUsersBoughtCourses(ctx context.Context, email string) ([]model.GetCourseLimitedResponse, error)
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

func (c CourseService) GetUserCourses(ctx context.Context, email string) ([]model.Course, error) {
	user, err := c.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	userCourses, err := c.courseRepo.GetUserCourses(user.Id)
	if err != nil {
		return nil, err
	}

	res, err := c.courseRepo.GetCoursesContent(ctx, userCourses)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c CourseService) GetUsersBoughtCourses(ctx context.Context, email string) ([]model.GetCourseLimitedResponse, error) {
	courses, err := c.courseRepo.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	user, err := c.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	var res []model.Course
	for _, ccrs := range courses {
		hasAccess, err := c.CheckCourseAccess(user.Id, ccrs.ID.Hex())
		if err != nil {
			return nil, errors.New("err getting access")
		}
		if !hasAccess {
			return nil, errors.New("course access if forbidden")
			//	return c.JSON(http.StatusForbidden, echo.Map{"error": "Course access if forbidden"})
		}
		res = append(res, ccrs)
	}

	var resp []model.GetCourseLimitedResponse
	for _, r := range res {
		resp = append(resp, r.ToRespModel())
	}

	return resp, nil
}
