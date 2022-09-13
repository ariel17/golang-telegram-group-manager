package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
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
	bh.Handle(debugHandler, th.CommandEqual(config.Debug))
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
		sentry.CaptureException(err)
	}
}

func welcomeHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID), services.GetWelcome(update.Message.Chat.ID),
	))
	if err != nil {
		sentry.CaptureException(err)
	}
}

func setWelcomeHandler(bot *telego.Bot, update telego.Update) {
	text := removeCommandFromText(update.Message.Text, config.SetWelcome)
	services.SetWelcome(text, update.Message.Chat.ID)
	_, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID), "Welcome message updated 🙌🏽",
	))
	if err != nil {
		sentry.CaptureException(err)
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
		inactives := services.GetInactives(days, update.Message.Chat.ID)
		text = services.FormatInactivesMessage("😴 Inactive users:\n", inactives)
	}

	_, err = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), text))
	if err != nil {
		sentry.CaptureException(err)
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
		inactives, untilDate, err := services.KickInactives(days, bot, update)
		if err != nil {
			sentry.CaptureException(err)
			text = errorToText(err)
		} else {
			text = services.FormatInactivesMessage("Users kicked 👋💔:\n", inactives)
			text += fmt.Sprintf(
				"\nThey are unable to re-join the group until %s\n🤷 Sorry-not sorry\n",
				untilDate.Format("2006-01-02 15:04"),
			)
		}
	}

	_, err = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), text))
	if err != nil {
		sentry.CaptureException(err)
	}
}

func activityHandler(bot *telego.Bot, update telego.Update) {
	if update.Message.From.IsBot {
		return
	}
	// TODO this is a workaround until I figure out what happens here and why it
	//  is not detected as command for multiline messages
	if strings.HasPrefix(update.Message.Text, fmt.Sprintf("/%s", config.SetWelcome)) {
		setWelcomeHandler(bot, update)
		return
	}
	services.SetActivityForUser(*update.Message)
}

func defaultHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), "I don't know this command 🤷🏽"),
	)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func statsHandler(bot *telego.Bot, update telego.Update) {
	_, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID),
			services.GetStatistics(update.Message.Chat.ID)),
	)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func debugHandler(bot *telego.Bot, update telego.Update) {
	text := removeCommandFromText(update.Message.Text, config.Debug)
	v, err := services.Debug(text)
	if err != nil {
		sentry.CaptureException(err)
		text = errorToText(err)
	} else if v == "" {
		text = "👍🖖"
	} else {
		text = v
	}
	_, err = bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), text),
	)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func errorToText(err error) string {
	return fmt.Sprintf("Can't complete that 🤔 The problem was: %v", err)
}

func removeCommandFromText(text, command string) string {
	v := strings.ReplaceAll(text, fmt.Sprintf("/%s ", command), "")
	return strings.ReplaceAll(v, fmt.Sprintf("/%s", command), "")
}