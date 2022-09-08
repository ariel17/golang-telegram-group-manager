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
	helpDescription = "Shows command usage."
)

var (
	descriptions map[string]string
)

// GetDescriptions returns a map of descriptions on existing commands.
func GetDescriptions() map[string]string {
	return descriptions
}

func init() {
	descriptions = map[string]string{
		Start:         helpDescription,
		Help:          helpDescription,
		Inactives:     fmt.Sprintf("Returns the list of inactive users in a time period. Usage: /%s 30d", Inactives),
		KickInactives: "This is the kick inactive command",
		Welcome:       "Shows the welcome message.",
		SetWelcome:    fmt.Sprintf("Saves a new welcome message. Usage: /%s <text>", SetWelcome),
	}
}