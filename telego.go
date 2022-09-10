package main

import (
	"log"

	"github.com/mymmrac/telego"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/handlers"
)

func main() {
	bot, err := telego.NewBot(config.GetTelegramApiToken(), telego.WithDefaultDebugLogger())
	if err != nil {
		panic(err)
	}

	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", me.Username)

	handlers.ConfigureHandlers(bot)
}