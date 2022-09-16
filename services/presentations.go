package services

import (
	"github.com/mymmrac/telego"

	"github.com/ariel17/golang-telegram-group-manager/config"
	"github.com/ariel17/golang-telegram-group-manager/repositories"
)

func GetPresentation(chatID, userID int64) (string, string, bool) {
	lang := repository.GetLangForChat(chatID)
	presentation, found := repository.GetPresentationForUser(chatID, userID)
	if !found {
		return config.GetNoPresentationText(lang), "", false
	}
	return presentation.Text, presentation.PhotoID, true
}

func SetPresentation(chatID, userID int64, text string, photos []telego.PhotoSize) string {
	lang := repository.GetLangForChat(chatID)
	if photos == nil {
		return config.GetMissingPhotoText(lang)
	}
	photo := photos[len(photos)-1]
	repository.SetPresentationForUser(chatID, userID, repositories.UserPresentation{
		Text:    text,
		PhotoID: photo.FileID,
	})
	return config.GetPresentationChangedText(lang)
}