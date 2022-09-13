package config

import (
	"fmt"
)

const (
	Start           = "start"
	Help            = "help"
	Inactives       = "inactives"
	KickInactives   = "kickinactives"
	Welcome         = "welcome"
	SetWelcome      = "setwelcome"
	Stats           = "stats"
	Debug           = "debug"
	helpDescription = "Shows command usage."
)

var (
	descriptions             map[string]string
	inactivesDescription     = fmt.Sprintf("Returns the list of inactive users in days period. Usage: /%s <days>", Inactives)
	kickInactivesDescription = fmt.Sprintf("Removes all inactive users from group in a time period. Usage: /%s <days>", KickInactives)
	setWelcomeDescription    = fmt.Sprintf("Saves a new welcome message. Usage: /%s <text>", SetWelcome)
)

// GetDescriptions returns a map of descriptions on existing commands.
func GetDescriptions() map[string]string {
	return descriptions
}

func init() {
	descriptions = map[string]string{
		Start:         helpDescription,
		Help:          helpDescription,
		Inactives:     inactivesDescription,
		KickInactives: kickInactivesDescription,
		Welcome:       "Shows the welcome message.",
		SetWelcome:    setWelcomeDescription,
		Stats:         "Shows user stats",
	}
}