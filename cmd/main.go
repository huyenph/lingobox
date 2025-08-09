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
	"github.com/huyenph/lingobox/model"
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

	config.DB.AutoMigrate(&model.User{}, &model.Word{}, &model.Example{})

	// Initialize bot with telebot
	if cfg.TelegramBotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	// Create new telebot instance
	telegramBot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.SetupHandlers(telegramBot)

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
