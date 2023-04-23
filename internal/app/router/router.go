package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ilyasoloviev99/quest_game_telegram_bot/internal/app/commands"
)

type Commander interface {
	HandleCommand(msg *tgbotapi.Message)
}

type Router struct {
	bot       *tgbotapi.BotAPI
	commander Commander
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:       bot,
		commander: commands.NewCommander(bot),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	r.commander.HandleCommand(update.Message)
}
