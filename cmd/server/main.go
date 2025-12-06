package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/db"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/routes"
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

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Register routes
	routes.RegisterRoutes(app, postgres, cfg)

	// Start server
	log.Printf("Starting server on port %s...", cfg.AppPort)
	app.Listen(":" + cfg.AppPort) // <-- capital L
}
