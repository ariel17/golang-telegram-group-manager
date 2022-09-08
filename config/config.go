package config

import (
	"errors"
	"os"
)

var telegramApiToken string

// GetTelegramApiToken returns the configured value of Telegram API token.
func GetTelegramApiToken() string {
	return telegramApiToken
}

func init() {
	telegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	if telegramApiToken == "" {
		panic(errors.New("telegram api token is required"))
	}

}