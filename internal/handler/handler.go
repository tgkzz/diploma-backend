package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"server/internal/service"
)

type Handler struct {
	service     *service.Service
	infoLogger  *log.Logger
	errorLogger *log.Logger
	jwtSecret   string
}

func NewHandler(service *service.Service, info *log.Logger, err *log.Logger, jwtSecret string) *Handler {
	return &Handler{
		service:     service,
		infoLogger:  info,
		errorLogger: err,
		jwtSecret:   jwtSecret,
	}
}

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	// CORS settings
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Rate limiter
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "ping")
	})

	authApi := e.Group("/auth")
	{
		authApi.POST("/send-email", h.sendEmailCode)
		authApi.POST("/check-code", h.checkEmailCode)
		authApi.POST("/register", h.register)
		authApi.POST("/login", h.login)
	}

	user := e.Group("/user")
	user.Use(h.jwtMiddleware)
	{
		user.GET("", h.getUserByEmail)
		user.DELETE("", h.deleteUser)
		user.PUT("", h.UpdateUser)

		user.GET("/get-courses", h.getUserCourses)
		user.GET("/get-meets", h.getUserMeets)

		user.POST("/buy-course/:course_id", h.buyCourse)
		userCourse := user.Group("")
		userCourse.Use(h.courseAccessMiddleware)
		{
			userCourse.GET("/:course_id", h.getUserCourse)
		}

		meeting := user.Group("/meeting")
		{
			meeting.POST("/make-appointment", h.makeAppointment)
			meeting.GET("/by-room-id", h.GetMeetingByRoomId)
		}
	}

	course := e.Group("/course")
	{
		course.GET("/get-all-courses", h.GetAllCourses)
		course.GET("/:course_id", h.GetCourse)
	}

	adminApi := e.Group("/admin")
	{
		adminApi.POST("/register", h.registerAdmin)
		adminApi.POST("/login", h.loginAdmin)
	}

	expertApi := e.Group("/expert")
	{
		expertApi.GET("/getAllExperts", h.getAllExperts)

		expertApi.POST("/register", h.registerExpert)
		expertApi.POST("/login", h.loginExpert)

		expertAction := expertApi.Group("")
		expertAction.Use(h.expertJwtMiddleware)
		{
			expertAction.POST("/create/meet", h.createMeet)
			expertAction.GET("/get-meets", h.GetExpertMeets)
		}

		expertApi.GET("/by-room-id", h.GetMeetingByRoomId)
	}

	chatApi := e.Group("/chat")
	{
		chatApi.GET("/ws/:room_id", h.handleConnections)
	}

	return e
}
