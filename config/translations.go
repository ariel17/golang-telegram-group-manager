package config

import "fmt"

const (
	ENGLISH_LANG = "en"
	SPANISH_LANG = "es"

	enHelpDescription = "Shows command usage."
	esHelpDescription = "Muestra el uso de los comandos."
)

var (
	descriptions = map[string]map[string]string{
		ENGLISH_LANG: {
			Start:         enHelpDescription,
			Help:          enHelpDescription,
			Inactives:     fmt.Sprintf("Returns the list of inactive users in days period. Usage: /%s <days>", Inactives),
			KickInactives: fmt.Sprintf("Removes all inactive users from group in a time period. Usage: /%s <days>", KickInactives),
			Welcome:       "Shows the welcome message.",
			SetWelcome:    fmt.Sprintf("Saves a new welcome message. Usage: /%s <text>", SetWelcome),
			Stats:         "Shows user stats",
			SetLang:       fmt.Sprintf("Sets the language. Usage: /%s <lang> (en: english, es: spanish)", SetLang),
		},
		SPANISH_LANG: {
			Start:         esHelpDescription,
			Help:          esHelpDescription,
			Inactives:     fmt.Sprintf("Retorna la lista de usuarios inactivos en un período de días. Uso: /%s <dias>", Inactives),
			KickInactives: fmt.Sprintf("Remueve todos los usuarios inactivos del grupo en un período de días. Uso: /%s <dias>", KickInactives),
			Welcome:       "Muestra el mensaje de bienvenida.",
			SetWelcome:    fmt.Sprintf("Guarda un nuevo mensaje de bienvenida. Uso: /%s <texto>", SetWelcome),
			Stats:         "Muestra las estadísticas de los usuarios.",
			SetLang:       fmt.Sprintf("Configura el lenguage. Uso: /%s <idioma> (en: inglés, es: español)", SetLang),
		},
	}
	availableCommands = map[string]string{
		ENGLISH_LANG: "🕹 Available commands:\n",
		SPANISH_LANG: "🕹 Comandos disponibles:\n",
	}
	languageError = map[string]string{
		ENGLISH_LANG: "That's not a valid language 🤷 Try with \"en\" or \"es\".",
		SPANISH_LANG: "Ese no es un idioma válido 🤷 Pruebe con \"en\" o \"es\".",
	}
	languageChanged = map[string]string{
		ENGLISH_LANG: "Language changed to english 🙌🏽",
		SPANISH_LANG: "Idioma cambiado a español 🙌🏽",
	}
	welcomeChanged = map[string]string{
		ENGLISH_LANG: "Welcome message updated 🙌🏽",
		SPANISH_LANG: "Mensaje de bienvenida actualizado 🙌🏽",
	}
	welcomeEmpty = map[string]string{
		ENGLISH_LANG: fmt.Sprintf("Hello! 👋 Welcome message is empty. You need to set one with /%s <text>", SetWelcome),
		SPANISH_LANG: fmt.Sprintf("Hola! 👋 El mensaje de bienvenida no está configurado. Se puede hacer con /%s <texto>", SetWelcome),
	}
	noInactives = map[string]string{
		ENGLISH_LANG: "No inactive users 🙌🏽",
		SPANISH_LANG: "Sin usuarios inactivos 🙌🏽",
	}
	inactives = map[string]string{
		ENGLISH_LANG: "😴 Inactive users:\n",
		SPANISH_LANG: "😴 Usuarios inactivos:\n",
	}
	kickedPrefix = map[string]string{
		ENGLISH_LANG: "Users kicked 👋💔:\n",
		SPANISH_LANG: "Users kicked 👋💔:\n",
	}
	kickedSuffix = map[string]string{
		ENGLISH_LANG: "\nThey are unable to re-join the group until %s\n🤷 Sorry-not sorry\n",
		SPANISH_LANG: "\nNo podrán volver a unirse al grupo hasta %s\n🤷 Agua y ajo, vieja\n",
	}
	noStatistics = map[string]string{
		ENGLISH_LANG: "I don't have statistics yet 🤷",
		SPANISH_LANG: "No tengo estadísticas aún 🤷",
	}
	statistics = map[string]string{
		ENGLISH_LANG: "📈 User statistics:\n",
		SPANISH_LANG: "📈 Estadísticas de usuario:\n",
	}
	statisticsRow = map[string]string{
		ENGLISH_LANG: "* @%s: messages: %d, last seen on: %s\n",
		SPANISH_LANG: "* @%s: mensajes: %d, visto por última vez: %s\n",
	}
	errorText = map[string]string{
		ENGLISH_LANG: "Can't complete that 🤔 The problem was: %v",
		SPANISH_LANG: "No puedo completar eso 🤔 El problema fue: %v",
	}
)

// GetDescriptions returns a map of descriptions on existing commands.
func GetDescriptions(lang string) map[string]string {
	return descriptions[lang]
}

func GetAvailableCommandsText(lang string) string {
	return availableCommands[lang]
}

func GetLanguageErrorText(lang string) string {
	return languageError[lang]
}

func GetLanguageChangedText(lang string) string {
	return languageChanged[lang]
}

func GetWelcomeMessageUpdatedText(lang string) string {
	return welcomeChanged[lang]
}

func GetWelcomeEmptyText(lang string) string {
	return welcomeEmpty[lang]
}

func GetNoInactiveText(lang string) string {
	return noInactives[lang]
}

func GetInactivesText(lang string) string {
	return inactives[lang]
}

func GetKickedPrefix(lang string) string {
	return kickedPrefix[lang]
}

func GetKickedSuffix(lang string) string {
	return kickedSuffix[lang]
}

func GetNoStatisticsText(lang string) string {
	return noStatistics[lang]
}

func GetStatisticsText(lang string) string {
	return statistics[lang]
}

func GetStatisticsRowText(lang string) string {
	return statisticsRow[lang]
}

func GetErrorText(lang string) string {
	return errorText[lang]
}