package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
	"strings"
)

func (h *Handler) jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header is required"})
		}

		split := strings.Split(tokenString, " ")
		if len(split) != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Bad token"})
		}

		token, err := jwt.ParseWithClaims(split[1], &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(h.jwtSecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "could not parse key"})
		}

		if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
			c.Set("email", claims.Email)
		} else {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		return next(c)
	}
}

func (h *Handler) expertJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header is required"})
		}

		split := strings.Split(tokenString, " ")
		if len(split) != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Bad token"})
		}

		token, err := jwt.ParseWithClaims(split[1], &model.JwtExpertClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(h.jwtSecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "could not parse key"})
		}

		if claims, ok := token.Claims.(*model.JwtExpertClaims); ok && token.Valid {
			c.Set("email", claims.Email)
		} else {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		return next(c)
	}
}

func (h *Handler) courseAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Get("email").(string)

		courseID := c.Param("course_id")

		user, err := h.service.Auth.GetUserByEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not get user info"})
		}

		hasAccess, err := h.service.Course.CheckCourseAccess(user.Id, courseID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "err getting access"})
		}
		if !hasAccess {
			return c.JSON(http.StatusForbidden, echo.Map{"error": "Course access if forbidden"})
		}

		return next(c)
	}
}
