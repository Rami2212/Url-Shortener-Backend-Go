package db

import (
	"context"
	"log"

	"github.com/Rami2212/Url-Shortener-Backend-Go/internal/config"
)

func ConnectRedis(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})

	// test redis health
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect Redis:", err)
	}

	log.Println("Connected to Redis")
	return rdb
}
