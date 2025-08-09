package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gopkg.in/tucnak/telebot.v2"

	"github.com/huyenph/lingobox/bot"
	"github.com/huyenph/lingobox/config"
	"github.com/huyenph/lingobox/models"
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

	// Auto migrate User and Word models
	config.DB.AutoMigrate(&models.User{}, &models.Word{}, &models.Example{})

	// Initialize bot with telebot
	if cfg.TelegramBotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	// Create new telebot instance
	telegramBot, err := telebot.NewBot(telebot.Settings{
		Token:cfg.TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Setup your bot handlers in a separate function (see below)
	bot.SetupHandlers(telegramBot)

	// Run the bot in a separate goroutine so it doesn't block Fiber
	go telegramBot.Start()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ LingoBox — healthy")
	})

	log.Printf("✅ Starting server on :%d\n", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
	log.Fatal(err)
	}
}
