package routes

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/handlers"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/repositories"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/services"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	urlRepo := repositories.NewURLRepository(db)
	urlService := services.NewURLService(urlRepo, cfg)

	urlHandler := handlers.NewURLHandler(urlService, cfg)
	healthHandler := handlers.NewHealthHandler()

	app.Get("/health", healthHandler.HealthCheck)

	api := app.Group("/api")
	api.Post("/shorten", urlHandler.Shorten)

	app.Get("/:code", urlHandler.Redirect)
}
