package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rifqi142/indico-be/internal/config"
	"github.com/rifqi142/indico-be/internal/controllers"
	"github.com/rifqi142/indico-be/internal/repository"
	"github.com/rifqi142/indico-be/internal/routes"
	"github.com/rifqi142/indico-be/internal/seeders"
	"github.com/rifqi142/indico-be/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	if err := config.InitDatabase(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Get database instance
	db := config.GetDB()

	// Run auto migration
	if err := config.RunAutoMigration(db); err != nil {
		log.Fatalf("Failed to run auto migration: %v", err)
	}

	// Run seeders (only in development)
	if cfg.AppEnv == "development" {
		seeders.RunAllSeeders(db)
	}

	// Parse JWT expiration
	jwtExpiration, err := time.ParseDuration(cfg.JWTExpiration)
	if err != nil {
		log.Fatalf("Invalid JWT expiration format: %v", err)
	}

	// Initialize repositories
	voucherRepo := repository.NewVoucherRepository(db)

	// Initialize services
	authService := services.NewAuthService(cfg.JWTSecret, jwtExpiration)
	voucherService := services.NewVoucherService(voucherRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	voucherController := controllers.NewVoucherController(voucherService)

	// Setup Gin
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, authController, voucherController, cfg.JWTSecret)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server starting on %s (Environment: %s)", addr, cfg.AppEnv)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
