package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
)

func (h *Handler) makeAppointment(c echo.Context) error {
	userEmail := c.Get("email")

	var req model.MakeAppointmentRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.Meeting.PlaceAppointment(req, userEmail.(string)); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrTimeInPast) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new meeting",
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetMeetingByRoomId(c echo.Context) error {

	roomId := c.QueryParam("room_id")

	meet, err := h.service.Meeting.GetMeetingByRoomId(roomId)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, meet)
}
