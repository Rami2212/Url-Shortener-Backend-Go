package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Rami2212/Url-Shortener-Backend-Go/db"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/routes"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load config
	cfg := config.LoadConfig()

	// Connect to DB
	postgres := db.ConnectPostgres(cfg)

	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app, postgres, redis, cfg)

	// Start server
	log.Printf("Starting server on port %s...", cfg.AppPort)
	app.listen(":" + cfg.AppPort)
}
