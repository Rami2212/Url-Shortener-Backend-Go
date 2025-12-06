package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/handlers"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/repositories"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/services"
)

func RegisterRoutes(
	app *fiber.App,
	db *gorm.DB,
	redis *redis.Client,
	cfg *config.Config,
) {
	// Dependencies
	urlRepo := repositories.NewURLRepository(db)
	urlService := services.NewURLService(urlRepo, redis, cfg)
	urlHandler := handlers.NewURLHandler(urlService, cfg)
	healthHandler := handlers.NewHealthHandler()

	// Health check
	app.Get("/health", healthHandler.HealthCheck)

	// API group
	api := app.Group("/api")
	api.Post("/shorten", urlHandler.Shorten)

	// Redirect route (must be after /api routes)
	app.Get("/:code", urlHandler.Redirect)
}
