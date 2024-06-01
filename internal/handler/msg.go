package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
)

func (h *Handler) sendMsg(c echo.Context) error {
	var req model.Msg
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Auth.SendMsg(req.To, req.Msg); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
