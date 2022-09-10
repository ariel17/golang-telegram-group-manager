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
	helpDescription = "Shows command usage."
)

var (
	descriptions  map[string]string
	inactives     = fmt.Sprintf("Returns the list of inactive users in days period. Usage: /%s 30", Inactives)
	kickInactives = fmt.Sprintf("Removes all inactive users from group in a time period. Usage: /%s 30", KickInactives)
	setWelcome    = fmt.Sprintf("Saves a new welcome message. Usage: /%s <text>", SetWelcome)
)

// GetDescriptions returns a map of descriptions on existing commands.
func GetDescriptions() map[string]string {
	return descriptions
}

func init() {
	descriptions = map[string]string{
		Start:         helpDescription,
		Help:          helpDescription,
		Inactives:     inactives,
		KickInactives: kickInactives,
		Welcome:       "Shows the welcome message.",
		SetWelcome:    setWelcome,
		Stats:         "Show user stats",
	}
}