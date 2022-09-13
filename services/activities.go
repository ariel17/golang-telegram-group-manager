package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

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

func FormatInactivesMessage(title string, inactives []repositories.UserActivity) string {
	if len(inactives) == 0 {
		return "No inactive users 🙌🏽"
	}
	text := title
	for _, activity := range inactives {
		lastSeen := activity.LastSeen.Format("2006-01-02 15:04")
		text += fmt.Sprintf("* @%s: %s\n", activity.Username, lastSeen)
	}
	return text
}

// KickInactives removes the inactive users and returns the list of them and the
// date until when they will be able to rejoin.
func KickInactives(days int, bot *telego.Bot, update telego.Update) ([]repositories.UserActivity, time.Time, error) {
	inactives := GetInactives(days, update.Message.Chat.ID)
	untilDate := time.Now().AddDate(0, 1, 0)
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
	}
	return inactives, untilDate, nil
}

// GetStatistics returns the amount of messages sent by user and last seen time.
func GetStatistics(chatID int64) string {
	activities := repository.GetActivities(chatID)
	if len(activities) == 0 {
		return "I don't have statistics yet 🤷"
	}
	text := "📈 User statistics:\n"
	for _, activity := range activities {
		lastSeen := activity.LastSeen.Format("2006-01-02 15:04")
		text += fmt.Sprintf("* @%s: messages: %d, last seen on: %s\n", activity.Username, activity.Count, lastSeen)
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

func init() {
	repository = repositories.New()
}