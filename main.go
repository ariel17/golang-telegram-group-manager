package main

import (
	"encoding/json"
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/services"
)

func main() {
	bot, err := tg.NewBotAPI(config.GetTelegramApiToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		b, _ := json.Marshal(update)
		log.Printf("New message: %s", b)

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			services.SetActivityForUser(*update.Message)
			continue
		}

		var (
			text    string
			command = update.Message.Command()
		)
		log.Printf("Command is: %s", command)

		switch command {
		case config.Help,
			config.Start:
			text = services.GetHelpMessage()
		case config.Inactives:
			inactives, err := services.GetInactives(update.Message.Text, config.Inactives)
			if err != nil {
				text = errorToText(err)
			} else {
				text = services.FormatInactivesMessage("😴 Inactive users:\n", inactives)
			}
		case config.KickInactives:
			inactives, err := services.KickInactives(update.Message.Text)
			if err != nil {
				text = errorToText(err)
			} else {
				text = services.FormatInactivesMessage("👋💔 Kicked users:\n", inactives)
			}
		case config.Welcome:
			text = services.GetWelcome()
		case config.SetWelcome:
			services.SetWelcome(update.Message.Text)
			text = "Welcome message updated 🙌🏽"
		default:
			text = "I don't know this command 🤷🏽"
		}
		log.Printf("Text to send is: %s", text)

		msg := tg.NewMessage(update.Message.Chat.ID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func errorToText(err error) string {
	return fmt.Sprintf("Could not complete that 🤔 The problem was: %v", err)
}