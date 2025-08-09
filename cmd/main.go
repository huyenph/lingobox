package main

import (
	"log"
	"time"

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

	telegramBot.Start()
}
