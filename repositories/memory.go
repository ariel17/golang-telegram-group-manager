package repositories

import (
	"encoding/json"

	"github.com/getsentry/sentry-go"
)

type chat struct {
	Welcome    string                 `json:"welcome"`
	Activities map[int64]UserActivity `json:"activities"`
}

type memoryRepository map[int64]chat

func (m memoryRepository) GetActivityForUser(chatID, userID int64) (UserActivity, bool) {
	c, exists := m[chatID]
	if !exists {
		return UserActivity{}, false
	}
	if c.Activities == nil {
		return UserActivity{}, false
	}
	activity, exists := c.Activities[userID]
	if !exists {
		return UserActivity{}, false
	}
	return activity, true
}

func (m memoryRepository) SetActivityForUser(chatID, userID int64, activity UserActivity) {
	c, exists := m[chatID]
	if !exists {
		c = chat{}
	}
	if c.Activities == nil {
		c.Activities = map[int64]UserActivity{}
	}
	c.Activities[userID] = activity
	m[chatID] = c
}

func (m memoryRepository) GetActivities(chatID int64) []UserActivity {
	c, exists := m[chatID]
	if !exists {
		return []UserActivity{}
	}
	if c.Activities == nil {
		return []UserActivity{}
	}
	l := []UserActivity{}
	for _, a := range c.Activities {
		l = append(l, a)
	}
	return l
}

func (m memoryRepository) GetWelcomeForChat(chatID int64) (string, bool) {
	c, exists := m[chatID]
	if !exists {
		return "", false
	}
	return c.Welcome, c.Welcome != ""
}

func (m memoryRepository) SetWelcomeForChat(chatID int64, text string) {
	c, exists := m[chatID]
	if !exists {
		c = chat{}
	}
	c.Welcome = text
	m[chatID] = c
}

func (m memoryRepository) Set(value string) error {
	defer sentry.Recover()

	var temp memoryRepository
	err := json.Unmarshal([]byte(value), &temp)
	if err != nil {
		return err
	}
	for k, v := range temp {
		m[k] = v
	}
	for k := range m {
		_, exists := temp[k]
		if !exists {
			delete(m, k)
		}
	}
	return nil
}

func (m memoryRepository) Dump() string {
	b, _ := json.Marshal(m)
	return string(b)
}