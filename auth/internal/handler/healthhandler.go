package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
