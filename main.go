package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ LingoBox — healthy")
	})

	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
