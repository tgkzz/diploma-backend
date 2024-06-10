package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/model"
	"strconv"
)

func (h *Handler) registerExpert(c echo.Context) error {
	var expert model.Expert
	if err := c.Bind(&expert); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	if err := h.service.ExpertAuth.CreateExpert(expert); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	successResponse := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created new expert",
	}
	h.infoLogger.Print("Successfully created new expert")
	return c.JSON(http.StatusCreated, successResponse)
}

func (h *Handler) loginExpert(c echo.Context) error {
	var creds model.Expert
	if err := c.Bind(&creds); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	exp, err := h.service.ExpertAuth.CheckExpertCreds(creds)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrIncorrectUsernameOrPassword) {
			return ErrorHandler(c, err, http.StatusNotFound)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	token, err := h.service.ExpertAuth.JwtExpertAuthorization(exp)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully logged as expert",
		"token":   token,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) getAllExperts(c echo.Context) error {
	res, err := h.service.ExpertAuth.GetAllExperts()
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully got all experts",
		"experts": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) createMeet(c echo.Context) error {
	expertEmail := c.Get("email")

	var req model.MakeAppointmentRequest
	if err := c.Bind(&req); err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}
	req.ExpertEmail = expertEmail.(string)

	res, err := h.service.Meeting.CreateMeetByExpert(c.Request().Context(), req)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrTimeInPast) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Successfully created meet by expert",
		"roomId":  res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetExpertMeets(c echo.Context) error {
	expertEmail := c.Get("email")

	res, err := h.service.Meeting.GetExpertMeets(expertEmail.(string))
	if err != nil {
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetExpertById(c echo.Context) error {
	expertId := c.Param("expert_id")

	id, err := strconv.Atoi(expertId)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	res, err := h.service.ExpertAuth.GetExpertById(id)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrNotFound) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status": "success",
		"expert": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) getExpertAvailableMeets(c echo.Context) error {
	expertId := c.Param("expert_id")

	id, err := strconv.Atoi(expertId)
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	res, err := h.service.Meeting.GetExpertAvailableMeets(id)
	if err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrNotFound) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status": "success",
		"expert": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) getExpert(c echo.Context) error {
	email := c.Get("email")

	res, err := h.service.ExpertAuth.GetExpertByEmail(email.(string))
	if err != nil {
		h.errorLogger.Print(err)
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"status": "success",
		"expert": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteMeetByExpert(c echo.Context) error {
	roomId := c.Param("room_id")

	if err := h.service.Meeting.DeleteMeetByExpert(roomId); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrImpossibleOperation) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteMeetByUser(c echo.Context) error {
	roomId := c.Param("room_id")

	if err := h.service.Meeting.DeleteMeetByUser(roomId); err != nil {
		h.errorLogger.Print(err)
		if errors.Is(err, model.ErrImpossibleOperation) {
			return ErrorHandler(c, err, http.StatusBadRequest)
		}
		return ErrorHandler(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
