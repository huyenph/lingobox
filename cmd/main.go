package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/huyenph/lingobox/config"
)

func main() {
	// load .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	// Load app configuration
	cfg := config.LoadConfig()

	// Connect to database
	config.Connect(cfg)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ LingoBox — healthy")
	})

	log.Printf("✅ Starting server on :%d\n", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
	log.Fatal(err)
	}
}
