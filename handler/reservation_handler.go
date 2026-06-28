package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/middleware"
	"github.com/SkillexSJ/SpotSync/service"
	"github.com/SkillexSJ/SpotSync/utils"
)

type ReservationHandler struct {
	reservationService *service.ReservationService
}

// create handler
func NewReservationHandler(reservationService *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{reservationService: reservationService}
}

// create reservation
func (h *ReservationHandler) CreateReservation(c *echo.Context) error {
	// get user id from jwt
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
			"Unauthorized",
			"Could not identify user",
		))
	}

	var req dto.CreateReservationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid request body",
			err.Error(),
		))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Validation failed",
			err.Error(),
		))
	}

	result, err := h.reservationService.CreateReservation(userID, req)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse(
				"Zone not found",
				err.Error(),
			))
		}
		if errors.Is(err, utils.ErrZoneFull) {
			return c.JSON(http.StatusConflict, dto.ErrorResponse(
				"Zone is full",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse(
		"Reservation confirmed successfully",
		result,
	))
}

// get my reservations
func (h *ReservationHandler) GetMyReservations(c *echo.Context) error {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
			"Unauthorized",
			"Could not identify user",
		))
	}

	result, err := h.reservationService.GetMyReservations(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"My reservations retrieved successfully",
		result,
	))
}

// cancel reservation
func (h *ReservationHandler) CancelReservation(c *echo.Context) error {
	userID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
			"Unauthorized",
			"Could not identify user",
		))
	}

	// get reservation id
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid reservation ID",
			"Reservation ID must be a valid number",
		))
	}

	if err := h.reservationService.CancelReservation(uint(id), userID); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse(
				"Reservation not found",
				err.Error(),
			))
		}
		if errors.Is(err, utils.ErrForbidden) {
			return c.JSON(http.StatusForbidden, dto.ErrorResponse(
				"Forbidden",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Reservation cancelled successfully",
		nil,
	))
}

// get all reservations
func (h *ReservationHandler) GetAllReservations(c *echo.Context) error {
	result, err := h.reservationService.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"All reservations retrieved successfully",
		result,
	))
}
