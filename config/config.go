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
	AppEnv           string
}

// getEnv returns environment variable or fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// LoadConfig loads configuration from environment variables or .env files (dev/staging)
func LoadConfig() *Config {
	// Determine APP_ENV
	appEnv := getEnv("APP_ENV", "dev")

	// Only load .env files for dev/staging
	if appEnv != "production" {
		envFile := ".env.dev"
		if appEnv == "stag" {
			envFile = ".env.stag"
		}
		if err := godotenv.Load(envFile); err != nil {
			log.Println("⚠️  No .env file found, continuing with environment variables")
		}
	}

	// Parse PORT
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
		AppEnv:           appEnv,
	}

	// Validate required fields
	if cfg.TelegramBotToken == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN is required")
	}
	if cfg.DatabaseURL == "" {
		log.Fatal("❌ DATABASE_URL is required")
	}

	return cfg
}
