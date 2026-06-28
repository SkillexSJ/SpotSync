package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/service"
	"github.com/SkillexSJ/SpotSync/utils"
)

type AuthHandler struct {
	authService *service.AuthService
}

// create handler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// user reg
func (h *AuthHandler) Register(c *echo.Context) error {
	var req dto.RegisterRequest

	// Bind JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid request body",
			err.Error(),
		))
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Validation failed",
			err.Error(),
		))
	}

	//  service layer
	result, err := h.authService.Register(req)
	if err != nil {
		if errors.Is(err, utils.ErrDuplicateEmail) {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
				"Registration failed",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse(
		"User registered successfully",
		result,
	))
}

// Login handler
func (h *AuthHandler) Login(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Invalid request body",
			err.Error(),
		))
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse(
			"Validation failed",
			err.Error(),
		))
	}

	//  service layer
	result, err := h.authService.Login(req)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
				"Login failed",
				err.Error(),
			))
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Internal server error",
			"Something went wrong",
		))
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse(
		"Login successful",
		result,
	))
}
