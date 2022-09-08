package services

import (
	"fmt"
	"strings"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

// GetHelpMessage shows the help text with commands usage.
func GetHelpMessage() string {
	text := "Available commands:\n"
	for k, v := range config.GetDescriptions() {
		text += fmt.Sprintf("/%s: %s\n", k, v)
	}
	return text
}

// IsCommand checks on the text message if it starts with a bot command.
func IsCommand(text string) bool {
	for k, _ := range config.GetDescriptions() {
		if strings.HasPrefix(text, fmt.Sprintf("/%s", k)) {
			return true
		}
	}
	return false
}