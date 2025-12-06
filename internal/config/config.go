package config

import "os"

type Config struct {
	AppPort string

	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB   string

	ReddisAddr string
	ReddisPass string
	ReddisDB   string

	BaseShortURL string
}

func LoadConfig() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),

		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:   os.Getenv("POSTGRES_DB"),

		ReddisAddr: os.Getenv("REDIS_ADDR"),
		ReddisPass: os.Getenv("REDIS_PASSWORD"),
		ReddisDB:   os.Getenv("REDIS_DB"),

		BaseShortURL: os.Getenv("BASE_SHORT_URL"),
	}
}
