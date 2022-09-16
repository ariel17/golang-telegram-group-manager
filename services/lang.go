package services

import "github.com/ariel17/golang-telegram-group-manager/config"

func SetLanguage(chatID int64, lang string) string {
	currentLang := repository.GetLangForChat(chatID)
	if lang != config.ENGLISH_LANG && lang != config.SPANISH_LANG {
		return config.GetLanguageErrorText(currentLang)
	}
	repository.SetLangForChat(chatID, lang)
	return config.GetLanguageChangedText(lang)
}