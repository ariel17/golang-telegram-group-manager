package services

import (
	"github.com/ariel17/golang-telegram-group-manager/config"
)

// SetWelcome saves the welcome message to show in the group.
func SetWelcome(text string, chatID int64) string {
	lang := repository.GetLangForChat(chatID)
	repository.SetWelcomeForChat(chatID, text)
	return config.GetWelcomeMessageUpdatedText(lang)
}

// GetWelcome returns the saved welcome message to show.
func GetWelcome(chatID int64) string {
	lang := repository.GetLangForChat(chatID)
	text, exists := repository.GetWelcomeForChat(chatID)
	if !exists {
		return config.GetWelcomeEmptyText(lang)
	}
	return text
}