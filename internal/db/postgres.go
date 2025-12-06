package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/models"
)

func ConnectPostgres(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDB, cfg.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connecting to PostgreSQL:", err)
	}

	// Run model migrations
	db.AutoMigrate(&models.URL{})

	log.Println("Connected to PostgreSQL")
	return db
}
