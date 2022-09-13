package services

import (
	"fmt"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

// GetHelpMessage shows the help text with commands usage.
func GetHelpMessage() string {
	text := "🕹 Available commands:\n"
	for k, v := range config.GetDescriptions() {
		text += fmt.Sprintf("* /%s: %s\n", k, v)
	}
	return text
}