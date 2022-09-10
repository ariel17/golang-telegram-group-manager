package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

	"github.com/ariel17/golang-telegram-group-manager/config"
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
		activities[message.From.ID] = UserActivity{
			ID:       message.From.ID,
			Username: message.From.Username,
			LastSeen: time.Unix(message.Date, 0),
			Count:    1,
		}
		return
	}
	v.LastSeen = time.Unix(message.Date, 0)
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
		if activity.LastSeen.Before(limit) {
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
		lastSeen := activity.LastSeen.Format("2006-01-02 15:04")
		text += fmt.Sprintf("* @%s: %s\n", activity.Username, lastSeen)
	}
	return text
}

// KickInactives removes the inactive users
func KickInactives(duration string) ([]UserActivity, error) {
	inactives, err := GetInactives(duration, config.KickInactives)
	if err != nil {
		return nil, err
	}
	for range inactives {
		// TODO kick user
	}
	return inactives, nil
}

func DebugHandler(bot *telego.Bot, update telego.Update) {
	b, _ := json.Marshal(activities)
	_, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), fmt.Sprintf("Activities: %s", b)),
	)
	if err != nil {
		panic(err)
	}
}