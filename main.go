package main

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/handlers"
)

func main() {
	defer sentry.Recover()

	bot, err := telego.NewBot(config.GetTelegramApiToken(), telego.WithDefaultDebugLogger())
	if err != nil {
		panic(err)
	}

	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", me.Username)

	updates, err := bot.UpdatesViaLongPulling(nil)
	if err != nil {
		panic(err)
	}
	defer bot.StopLongPulling()

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		panic(err)
	}
	handlers.ConfigureHandlers(bh)
}