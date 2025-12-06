package config

import "os"

type Config struct {
	AppPort string

	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB   string

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

		BaseShortURL: os.Getenv("BASE_SHORT_URL"),
	}
}
