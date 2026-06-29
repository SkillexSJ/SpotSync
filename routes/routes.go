package routes

import (
	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/handler"
	"github.com/SkillexSJ/SpotSync/middleware"
)

func SetupRoutes(
	e *echo.Echo,
	authHandler *handler.AuthHandler,
	zoneHandler *handler.ZoneHandler,
	reservationHandler *handler.ReservationHandler,
	jwtSecret string,
) {

	//Base API

	api := e.Group("/api/v1")

	// Auth Routes

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// Zones Routes

	zones := api.Group("/zones")
	zones.GET("", zoneHandler.GetAllZones)
	zones.GET("/:id", zoneHandler.GetZoneByID)

	// Admin-only zone routes
	adminZones := api.Group("/zones", middleware.JWTMiddleware(jwtSecret), middleware.RequireRole("admin"))
	adminZones.POST("", zoneHandler.CreateZone)
	adminZones.PUT("/:id", zoneHandler.UpdateZone)
	adminZones.DELETE("/:id", zoneHandler.DeleteZone)

	// Reservation Routes
	reservations := api.Group("/reservations", middleware.JWTMiddleware(jwtSecret))
	reservations.POST("", reservationHandler.CreateReservation)
	reservations.GET("/my-reservations", reservationHandler.GetMyReservations)
	reservations.DELETE("/:id", reservationHandler.CancelReservation)

	// Admin-only reservation route
	adminReservations := api.Group("/reservations", middleware.JWTMiddleware(jwtSecret), middleware.RequireRole("admin"))
	adminReservations.GET("", reservationHandler.GetAllReservations)
}
