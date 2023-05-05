package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ilyasoloviev99/quest_game_telegram_bot/internal/app/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	routerHandler := router.NewRouter(bot)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
