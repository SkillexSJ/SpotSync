package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/service"
	"github.com/SkillexSJ/SpotSync/utils"
)

type ZoneHandler struct {
	zoneService *service.ZoneService
}

// create handler
func NewZoneHandler(zoneService *service.ZoneService) *ZoneHandler {
	return &ZoneHandler{zoneService: zoneService}
}

// create zone
func (h *ZoneHandler) CreateZone(c *echo.Context) error {
	var req dto.CreateZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid request body",
			err.Error(),
		))
	}

	// validate
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Validation failed",
			err.Error(),
		))
	}

	// service layer
	result, err := h.zoneService.CreateZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse(
		"Parking zone created successfully",
		result,
	))
}

// get all public
func (h *ZoneHandler) GetAllZones(c *echo.Context) error {
	result, err := h.zoneService.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Parking zones retrieved successfully",
		result,
	))
}

// get by id
func (h *ZoneHandler) GetZoneByID(c *echo.Context) error {
	// Parse zone ID from URL param
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid zone ID",
			"Zone ID must be a valid number",
		))
	}

	// service layer
	result, err := h.zoneService.GetZoneByID(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse(
				"Zone not found",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Parking zone retrieved successfully",
		result,
	))
}

// update zone
func (h *ZoneHandler) UpdateZone(c *echo.Context) error {
	// get id param
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid zone ID",
			"Zone ID must be a valid number",
		))
	}

	var req dto.UpdateZoneRequest
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

	// service layer
	result, err := h.zoneService.UpdateZone(uint(id), req)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse(
				"Zone not found",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Parking zone updated successfully",
		result,
	))
}

// delete zone
func (h *ZoneHandler) DeleteZone(c *echo.Context) error {
	// get id param
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid zone ID",
			"Zone ID must be a valid number",
		))
	}

	if err := h.zoneService.DeleteZone(uint(id)); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse(
				"Zone not found",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Parking zone deleted successfully",
		nil,
	))
}
