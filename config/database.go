package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect(cfg *Config) {
	database, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("✅ Connected to PostgreSQL")
}

// AutoMigrate runs GORM's auto migration for models
func AutoMigrate(models ...interface{}) {
	if DB == nil {
		log.Fatal("❌ Database connection is not initialized")
	}

	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	log.Println("✅ Database migration completed")
}
