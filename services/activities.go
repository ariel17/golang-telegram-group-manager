package services

import (
	"fmt"
	"strconv"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

type UserActivity struct {
	Message tg.Message
	Count   int64 `json:"count"`
}

func (u UserActivity) LastSeen() time.Time {
	return time.Unix(int64(u.Message.Date), 0)
}

var (
	activities = map[int64]UserActivity{}
)

// SetActivityForUser saves the last sent message
func SetActivityForUser(message tg.Message) {
	v, exists := activities[message.From.ID]
	if !exists {
		activities[message.From.ID] = UserActivity{
			Message: message,
			Count:   1,
		}
		return
	}
	v.Message = message
	v.Count += 1
	activities[message.From.ID] = v
}

// GetInactives returns the list of users without activity for the indicated
// time delta.
func GetInactives(duration, command string) ([]UserActivity, error) {
	parsedDuration := removeCommandFromText(duration, command)
	days, err := strconv.Atoi(parsedDuration)
	if err != nil {
		return nil, err
	}
	limit := time.Now().AddDate(0, 0, -days)

	inactives := []UserActivity{}
	for _, activity := range activities {
		if activity.LastSeen().Before(limit) {
			inactives = append(inactives, activity)
		}
	}
	return inactives, nil
}

func FormatInactivesMessage(title string, inactives []UserActivity) string {
	if len(inactives) == 0 {
		return "No inactive users 🙌🏽"
	}

	text := title
	for _, activity := range inactives {
		user := activity.Message.From
		lastSeen := activity.LastSeen().Format("2006-01-02 15:04")
		text += fmt.Sprintf("* @%s: %s\n", user.UserName, lastSeen)
	}
	return text
}

// KickInactives removes the inactive users
func KickInactives(duration string, bot *tg.BotAPI, chat tg.Chat) ([]UserActivity, error) {
	inactives, err := GetInactives(duration, config.KickInactives)
	if err != nil {
		return nil, err
	}
	for _, user := range inactives {
		c := tg.KickChatMemberConfig{
			ChatMemberConfig: tg.ChatMemberConfig{
				ChatID: chat.ID,
				UserID: user.Message.From.ID,
			},
			UntilDate:      0,
			RevokeMessages: false,
		}
		tg.Edit
	}
	return inactives, nil
}