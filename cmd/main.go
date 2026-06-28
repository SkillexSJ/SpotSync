package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"

	"github.com/SkillexSJ/SpotSync/config"
	"github.com/SkillexSJ/SpotSync/handler"
	"github.com/SkillexSJ/SpotSync/repository"
	"github.com/SkillexSJ/SpotSync/routes"
	"github.com/SkillexSJ/SpotSync/service"
)

// ──────────────────────────────────────────
// Custom Validator Adapter
// ──────────────────────────────────────────

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func main() {

	// Environment Variables

	config.LoadEnv()
	jwtSecret := config.GetJWTSecret()

	// Connect to Database
	db := config.ConnectDatabase()

	pgDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB: ", err)
	}
	if err := pgDB.Ping(); err != nil {
		log.Fatal("Database ping failed: ", err)
	}
	fmt.Println("✅ Database ping successful!")

	// Initialize Echo & Custom Validator
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	zoneRepo := repository.NewZoneRepository(db)
	reservationRepo := repository.NewReservationRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, jwtSecret)
	zoneService := service.NewZoneService(zoneRepo)
	reservationService := service.NewReservationService(reservationRepo, zoneRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	zoneHandler := handler.NewZoneHandler(zoneService)
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Health Check
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "SpotSync API is running",
		})
	})

	// Register Routes
	routes.SetupRoutes(e, authHandler, zoneHandler, reservationHandler, jwtSecret)

	// Start Server
	port := config.GetPort()
	fmt.Println("──────────────────────────────────────────")
	fmt.Printf("🚀 SpotSync server running on port %s\n", port)
	fmt.Println("──────────────────────────────────────────")

	if err := e.Start(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
