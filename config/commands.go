package config

const (
	Start         = "start"
	Help          = "help"
	Inactives     = "inactives"
	KickInactives = "kickinactives"
	Welcome       = "welcome"
	Unknown       = "I don't know this command"
)

var descriptions map[string]string

// GetDescriptions returns a map of descriptions on existing commands.
func GetDescriptions() map[string]string {
	return descriptions
}

func init() {
	descriptions = map[string]string{
		Start:         "This is the start command",
		Help:          "This is the help command",
		Inactives:     "This is the inactive command",
		KickInactives: "This is the kick inactive command",
		Welcome:       "This is the welcome message",
	}
}