package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/services"
)

// ConfigureHandlers TODO
func ConfigureHandlers(bh *th.BotHandler) {
	bh.Handle(helpHandler, th.CommandEqual(config.Help))
	bh.Handle(helpHandler, th.CommandEqual(config.Start))
	bh.Handle(welcomeHandler, th.CommandEqual(config.Welcome))
	bh.Handle(setWelcomeHandler, th.CommandEqual(config.SetWelcome))
	bh.Handle(inactivesHandler, th.CommandEqual(config.Inactives))
	bh.Handle(kickInactivesHandler, th.CommandEqual(config.KickInactives))
	bh.Handle(statsHandler, th.CommandEqual(config.Stats))

	bh.Handle(defaultHandler, th.AnyCommand())
	bh.Handle(activityHandler, th.AnyMessage())

	defer bh.Stop()
	bh.Start()
}

func helpHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID), services.GetHelpMessage(),
	))
	if err != nil {
		panic(err)
	}
}

func welcomeHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID), services.GetWelcome(),
	))
	if err != nil {
		panic(err)
	}
}

func setWelcomeHandler(bot *telego.Bot, update telego.Update) {
	text := removeCommandFromText(update.Message.Text, config.SetWelcome)
	services.SetWelcome(text)
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID), "Welcome message updated 🙌🏽",
	))
	if err != nil {
		panic(err)
	}
}

func inactivesHandler(bot *telego.Bot, update telego.Update) {
	var (
		text      string
		duration  = removeCommandFromText(update.Message.Text, config.Inactives)
		days, err = strconv.Atoi(duration)
	)

	if err != nil {
		text = errorToText(err)
	} else {
		inactives := services.GetInactives(days)
		text = services.FormatInactivesMessage("😴 Inactive users:\n", inactives)
	}

	_, err = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), text))
	if err != nil {
		panic(err)
	}
}

func kickInactivesHandler(bot *telego.Bot, update telego.Update) {
	var (
		text      string
		duration  = removeCommandFromText(update.Message.Text, config.KickInactives)
		days, err = strconv.Atoi(duration)
	)

	if err != nil {
		text = errorToText(err)
	} else {
		inactives := services.KickInactives(days)
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

func statsHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), services.GetStatistics()),
	)
	if err != nil {
		panic(err)
	}
}

func errorToText(err error) string {
	return fmt.Sprintf("Could not complete that 🤔 The problem was: %v", err)
}

func removeCommandFromText(text, command string) string {
	return strings.ReplaceAll(text, fmt.Sprintf("/%s ", command), "")
}