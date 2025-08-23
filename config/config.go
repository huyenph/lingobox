package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             int
	TelegramBotToken string
	DatabaseURL      string
	AuthorUsername   string
	AuthorEmail      string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	envFile := ".env"
	if os.Getenv("APP_ENV") == "dev" {
		envFile = ".env.dev"
	} else if os.Getenv("APP_ENV") == "stag" {
		envFile = ".env.stag"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("No .env file found, reading config from environment")
	}

	port, err := strconv.Atoi(getEnv("PORT", "3000"))
	if err != nil {
		log.Fatalf("Invalid PORT value: %v", err)
	}
	cfg := &Config{
		Port:             port,
		TelegramBotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		AuthorUsername:   getEnv("AUTHOR_USERNAME", ""),
		AuthorEmail:      getEnv("AUTHOR_EMAIL", ""),
	}

	if cfg.TelegramBotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}
