package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/utils"
)

// Context keys
const (
	ContextKeyUserID = "user_id"
	ContextKeyRole   = "role"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// get auth header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
					"Missing authorization header",
					"Authorization header is required",
				))
			}

			// expect format
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
					"Invalid authorization format",
					"Authorization header must be: Bearer <token>",
				))
			}

			tokenString := parts[1]

			// validate token
			claims, err := utils.ValidateToken(tokenString, secret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
					"Invalid or expired token",
					err.Error(),
				))
			}

			// insert the user into context
			c.Set(ContextKeyUserID, claims.UserID)
			c.Set(ContextKeyRole, claims.Role)

			return next(c)
		}
	}
}

// require role
func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// get role from context
			role, ok := c.Get(ContextKeyRole).(string)
			if !ok || role == "" {
				return c.JSON(http.StatusUnauthorized, dto.ErrorResponse(
					"Unauthorized",
					"Could not determine user role",
				))
			}

			// check user allowed roles
			for _, allowed := range allowedRoles {
				if role == allowed {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, dto.ErrorResponse(
				"Forbidden",
				"You do not have permission to access this resource",
			))
		}
	}
}

// get user id from context
func GetUserIDFromContext(c *echo.Context) (uint, bool) {
	userID, ok := c.Get(ContextKeyUserID).(uint)
	return userID, ok
}

// get role from context
func GetRoleFromContext(c *echo.Context) (string, bool) {
	role, ok := c.Get(ContextKeyRole).(string)
	return role, ok
}
