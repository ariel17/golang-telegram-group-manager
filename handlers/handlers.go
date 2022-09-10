package handlers

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/services"
)

// ConfigureHandlers TODO
func ConfigureHandlers(bot *telego.Bot) {
	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	bh.Handle(helpHandler, th.CommandEqual(config.Help))
	bh.Handle(helpHandler, th.CommandEqual(config.Start))
	bh.Handle(welcomeHandler, th.CommandEqual(config.Welcome))
	bh.Handle(setWelcomeHandler, th.CommandEqual(config.SetWelcome))
	bh.Handle(inactivesHandler, th.CommandEqual(config.Inactives))
	bh.Handle(kickInactivesHandler, th.CommandEqual(config.KickInactives))
	// TODO stats handler

	bh.Handle(services.DebugHandler, th.CommandEqual("debug"))

	bh.Handle(defaultHandler, th.AnyCommand())
	bh.Handle(activityHandler, th.AnyMessage())
	bh.Start()
}

func helpHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		services.GetHelpMessage(),
	))
	if err != nil {
		panic(err)
	}
}

func welcomeHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		services.GetWelcome(),
	))
	if err != nil {
		panic(err)
	}
}

func setWelcomeHandler(bot *telego.Bot, update telego.Update) {
	services.SetWelcome(update.Message.Text)
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Welcome message updated 🙌🏽",
	))
	if err != nil {
		panic(err)
	}
}

func inactivesHandler(bot *telego.Bot, update telego.Update) {
	var text string
	inactives, err := services.GetInactives(update.Message.Text, config.Inactives)
	if err != nil {
		text = errorToText(err)
	} else {
		text = services.FormatInactivesMessage("😴 Inactive users:\n", inactives)
	}
	_, err = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), text))
	if err != nil {
		panic(err)
	}
}

func kickInactivesHandler(bot *telego.Bot, update telego.Update) {
	var text string
	inactives, err := services.KickInactives(update.Message.Text)
	if err != nil {
		text = errorToText(err)
	} else {
		text = services.FormatInactivesMessage("👋💔 Kicked users:\n", inactives)
	}
	_, err = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), text))
	if err != nil {
		panic(err)
	}
}

func activityHandler(_ *telego.Bot, update telego.Update) {
	services.SetActivityForUser(*update.Message)
}

func defaultHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), "I don't know this command 🤷🏽"),
	)
	if err != nil {
		panic(err)
	}
}

func errorToText(err error) string {
	return fmt.Sprintf("Could not complete that 🤔 The problem was: %v", err)
}