package services

import (
	"fmt"
	"time"

	"github.com/mymmrac/telego"
)

type UserActivity struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	LastSeen time.Time `json:"last_seen"`
	Count    int64     `json:"count"`
}

var (
	activities = map[int64]UserActivity{}
)

// SetActivityForUser saves the last sent message
func SetActivityForUser(message telego.Message) {
	v, exists := activities[message.From.ID]
	if !exists {
		v = UserActivity{
			ID:       message.From.ID,
			Username: message.From.Username,
		}
	}
	v.LastSeen = time.Unix(message.Date, 0)
	v.Count += 1
	activities[message.From.ID] = v
}

// GetInactives returns the list of users without activity for the indicated
// time delta.
func GetInactives(days int) []UserActivity {
	limit := time.Now().AddDate(0, 0, -days)
	inactives := []UserActivity{}
	for _, activity := range activities {
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
func KickInactives(days int) []UserActivity {
	inactives := GetInactives(days)
	for range inactives {
		// TODO kick user
	}
	return inactives
}

// GetStatistics returns the amount of messages sent by user and last seen time.
func GetStatistics() string {
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