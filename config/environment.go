package config

import (
	"os"
)

var (
	telegramApiToken string
	DebugJSON        string
)

// GetTelegramApiToken returns the configured value of Telegram API token.
func GetTelegramApiToken() string {
	return telegramApiToken
}

func init() {
	telegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	DebugJSON = os.Getenv("DEBUG_JSON")
}