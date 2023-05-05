package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander interface {
	HandleCommand(msg *tgbotapi.Message)
}

type DefaultCommander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI) *DefaultCommander {
	return &DefaultCommander{bot: bot}
}

func (c *DefaultCommander) HandleCommand(msg *tgbotapi.Message) {
	// TODO: Adds commands
	newMessage := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("You wrote: %s", msg.Text))

	_, err := c.bot.Send(newMessage)
	if err != nil {
		log.Printf("Error sending message")
	}
}
