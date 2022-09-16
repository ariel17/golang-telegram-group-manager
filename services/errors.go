package services

import (
	"fmt"

	"github.com/ariel17/golang-telegram-group-manager/config"
)

func ErrorToText(chatID int64, err error) string {
	lang := repository.GetLangForChat(chatID)
	return fmt.Sprintf(config.GetErrorText(lang), err)
}