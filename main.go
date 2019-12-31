package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

const botTokenKey = "TG_BOT_TOKEN"

func main() {
	botToken := loadConfig()
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error starting bot: %s", err.Error())
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err.Error())
	}
	for update := range updates {
		if update.InlineQuery != nil {
			processInlineQuery(update.InlineQuery, bot)
		}
	}
}

func loadConfig() string {
	log.Println("Loading config...")
	err := godotenv.Load()
	if err != nil {
		log.Println("no config found")
	}
	bt := os.Getenv(botTokenKey)
	if bt == "" {
		log.Fatalf("No bot token, please set %s. Exiting.", botTokenKey)
	}
	return bt
}
