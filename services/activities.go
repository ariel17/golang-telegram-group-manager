package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type UserActivity struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	LastSeen time.Time `json:"last_seen"`
	Count    int64     `json:"count"`
}

var (
	activities = map[int64]map[int64]UserActivity{}
)

// SetActivityForUser saves the last sent message
func SetActivityForUser(message telego.Message) {
	chat, exists := activities[message.Chat.ID]
	if !exists {
		chat = map[int64]UserActivity{}
	}

	activity, exists := chat[message.From.ID]
	if !exists {
		activity = UserActivity{
			ID:       message.From.ID,
			Username: determineUsername(*message.From),
		}
	}
	activity.LastSeen = time.Unix(message.Date, 0)
	activity.Count += 1
	chat[message.From.ID] = activity
	activities[message.Chat.ID] = chat
}

// GetInactives returns the list of users without activity for the indicated
// time delta.
func GetInactives(days int, chatID int64) []UserActivity {
	inactives := []UserActivity{}
	chat, exists := activities[chatID]
	if !exists {
		return inactives
	}

	limit := time.Now().AddDate(0, 0, -days)
	for _, activity := range chat {
		if activity.LastSeen.Before(limit) {
			inactives = append(inactives, activity)
		}
	}
	return inactives
}

func FormatInactivesMessage(title string, inactives []UserActivity) string {
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

// KickInactives removes the inactive users
func KickInactives(days int, bot *telego.Bot, update telego.Update) ([]UserActivity, error) {
	inactives := GetInactives(days, update.Message.Chat.ID)
	for _, inactive := range inactives {
		unbanParams := telego.UnbanChatMemberParams{
			ChatID: tu.ID(update.Message.Chat.ID),
			UserID: inactive.ID,
		}
		err := bot.UnbanChatMember(&unbanParams)
		if err != nil {
			return nil, err
		}
	}
	return inactives, nil
}

// GetStatistics returns the amount of messages sent by user and last seen time.
func GetStatistics(chatID int64) string {
	chat, exists := activities[chatID]
	if !exists {
		return "I don't have statistics yet 🤷"
	}
	text := "📈 User statistics:\n"
	for _, activity := range chat {
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