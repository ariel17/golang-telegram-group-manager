package services

import (
	"fmt"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

// GetHelpMessage shows the help text with commands usage.
func GetHelpMessage(chatID int64) string {
	lang := repository.GetLangForChat(chatID)
	text := config.GetAvailableCommandsText(lang)
	for k, v := range config.GetDescriptions(lang) {
		text += fmt.Sprintf("* /%s: %s\n", k, v)
	}
	return text
}