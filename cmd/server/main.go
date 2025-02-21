package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahar3/trip-planning-service/config"
	"github.com/shahar3/trip-planning-service/handlers"
	"github.com/shahar3/trip-planning-service/repository"
	"github.com/shahar3/trip-planning-service/routes"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize Gin router (you can add middleware here as needed)
	router := gin.Default()

	// Initialize in-memory repository and trip handler
	tripRepo := repository.NewInMemoryTripRepository()
	tripHandler := handlers.NewTripHandler(tripRepo)

	// Setup API routes
	routes.SetupRoutes(router, tripHandler)

	// Configure HTTP server settings
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Printf("Trip Planning Service running on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
