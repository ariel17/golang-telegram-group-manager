package services

import (
	"fmt"
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
func GetInactives(duration string) ([]UserActivity, error) {
	duration = removeCommandFromText(duration, config.Inactives)
	d, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	limit := time.Now().Add(-d)

	inactives := []UserActivity{}
	for _, activity := range activities {
		if activity.LastSeen().Before(limit) {
			inactives = append(inactives, activity)
		}
	}
	return inactives, nil
}

func FormatInactivesMessage(inactives []UserActivity) string {
	if len(inactives) == 0 {
		return "No inactive users :)"
	}

	text := "Inactive users:\n"
	for _, activity := range inactives {
		user := activity.Message.From
		text += fmt.Sprintf("* %s (%d): %v\n", user.UserName, user.ID, activity.LastSeen())
	}
	return text
}