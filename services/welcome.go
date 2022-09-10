package services

import (
	"fmt"
	"strings"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

var welcomeMessage string

// SetWelcome saves the welcome message to show in the group.
func SetWelcome(text string) {
	welcomeMessage = removeCommandFromText(text, config.SetWelcome)
}

// GetWelcome returns the saved welcome message to show.
func GetWelcome() string {
	if welcomeMessage == "" {
		return fmt.Sprintf("Hello! 👋 Welcome message is empty. You need to set one with /%s <text>", config.SetWelcome)
	}
	return welcomeMessage
}

func removeCommandFromText(text, command string) string {
	return strings.ReplaceAll(text, fmt.Sprintf("/%s ", command), "")
}