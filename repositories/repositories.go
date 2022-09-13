package repositories

import "time"

// UserActivity collects the basic activity for user in chat.
type UserActivity struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	LastSeen time.Time `json:"last_seen"`
	Count    int64     `json:"count"`
}

type Repository interface {
	GetActivityForUser(chatID, userID int64) (UserActivity, bool)
	SetActivityForUser(chatID, userID int64, activity UserActivity)
	GetActivities(chatID int64) ([]UserActivity, bool)
	GetWelcomeForChat(chatID int64) (string, bool)
	SetWelcomeForChat(chatID int64, text string)
	Set(value string) error
	Dump() string
}