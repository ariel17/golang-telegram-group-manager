package main

import (
	"encoding/json"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ariel17/golang-telegram-group-manager/config"
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

		if !update.Message.IsCommand() {
			continue
		}

		var (
			text         string
			descriptions = config.GetDescriptions()
			command      = update.Message.Command()
		)
		log.Printf("Command is: %s", command)

		switch command {
		case config.Help,
			config.Inactives,
			config.Start,
			config.Welcome,
			config.KickInactives:
			text = descriptions[command]
		default:
			text = descriptions[config.Unknown]
		}
		log.Printf("Text to send is: %s", text)

		msg := tg.NewMessage(update.Message.Chat.ID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}