package services

import (
	"fmt"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

var (
	welcomeMessage = map[int64]string{}
)

// SetWelcome saves the welcome message to show in the group.
func SetWelcome(text string, chatID int64) {
	welcomeMessage[chatID] = text
}

// GetWelcome returns the saved welcome message to show.
func GetWelcome(chatID int64) string {
	text, exists := welcomeMessage[chatID]
	if !exists {
		return fmt.Sprintf("Hello! 👋 Welcome message is empty. You need to set one with /%s <text>", config.SetWelcome)
	}
	return text
}