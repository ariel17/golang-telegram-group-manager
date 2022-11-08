package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/repositories"
)

var (
	repository repositories.Repository
)

// SetActivityForUser saves the last sent message
func SetActivityForUser(message telego.Message) {
	activity, exists := repository.GetActivityForUser(message.Chat.ID, message.From.ID)
	if !exists {
		activity = repositories.UserActivity{
			ID:       message.From.ID,
			Username: determineUsername(*message.From),
		}
	}
	activity.LastSeen = time.Unix(message.Date, 0)
	activity.Count += 1
	repository.SetActivityForUser(message.Chat.ID, message.From.ID, activity)
}

// GetInactives returns the list of users without activity for the indicated
// time delta.
func GetInactives(days int, chatID int64) []repositories.UserActivity {
	activities := repository.GetActivities(chatID)
	limit := time.Now().AddDate(0, 0, -days)
	inactives := []repositories.UserActivity{}
	for _, activity := range activities {
		if activity.LastSeen.Before(limit) {
			inactives = append(inactives, activity)
		}
	}
	return inactives
}

func FormatInactivesMessage(chatID int64, inactives []repositories.UserActivity) string {
	lang := repository.GetLangForChat(chatID)
	title := config.GetInactivesText(lang)
	text, _ := formatInactives(lang, title, inactives)
	return text
}

func FormatKickedInactivesMessage(chatID int64, inactives []repositories.UserActivity, untilDate time.Time) string {
	lang := repository.GetLangForChat(chatID)
	prefix := config.GetKickedPrefix(lang)
	text, formatted := formatInactives(lang, prefix, inactives)
	if !formatted {
		return text
	}
	suffix := fmt.Sprintf(config.GetKickedSuffix(lang), untilDate.Format("2006-01-02 15:04"))
	return text + suffix
}

// KickInactives removes the inactive users and returns the list of them and the
// date until when they will be able to rejoin.
func KickInactives(days int, bot *telego.Bot, update telego.Update) ([]repositories.UserActivity, time.Time, error) {
	inactives := GetInactives(days, update.Message.Chat.ID)
	untilDate := time.Now().AddDate(0, 0, days)
	for _, inactive := range inactives {
		banParams := telego.BanChatMemberParams{
			ChatID:         tu.ID(update.Message.Chat.ID),
			UserID:         inactive.ID,
			UntilDate:      untilDate.Unix(),
			RevokeMessages: false,
		}
		err := bot.BanChatMember(&banParams)
		if err != nil {
			return nil, time.Time{}, err
		}
		RemoveUser(update.Message.Chat.ID, inactive.ID)
	}
	return inactives, untilDate, nil
}

// GetStatistics returns the amount of messages sent by user and last seen time.
func GetStatistics(chatID int64) string {
	lang := repository.GetLangForChat(chatID)
	activities := repository.GetActivities(chatID)
	if len(activities) == 0 {
		return config.GetNoStatisticsText(lang)
	}
	text := config.GetStatisticsText(lang)
	for _, activity := range activities {
		lastSeen := activity.LastSeen.Format("2006-01-02 15:04")
		text += fmt.Sprintf(config.GetStatisticsRowText(lang), activity.Username, activity.Count, lastSeen)
	}
	return text
}

func determineUsername(user telego.User) string {
	if user.Username != "" {
		return user.Username
	}
	v := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	return strings.TrimSpace(v)
}

func formatInactives(lang, title string, inactives []repositories.UserActivity) (string, bool) {
	if len(inactives) == 0 {
		return config.GetNoInactiveText(lang), false
	}
	text := title
	for _, activity := range inactives {
		lastSeen := activity.LastSeen.Format("2006-01-02 15:04")
		text += fmt.Sprintf("* @%s: %s\n", activity.Username, lastSeen)
	}
	return text, true
}

func RemoveUser(chatID, userID int64) {
	repository.RemoveUserData(chatID, userID)
}

func init() {
	repository = repositories.New()
	if config.DebugJSON != "" {
		if err := repository.Set(config.DebugJSON); err != nil {
			log.Printf("Failed loading debug JSON: %v", err)
			sentry.CaptureException(err)
		} else {
			log.Printf("Succesful loading of debug JSON")
		}
		config.DebugJSON = ""
	}
}