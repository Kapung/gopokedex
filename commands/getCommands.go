package commands

type cliCommand struct {
	name        string
	description string
	Callback    func()
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Opens up the map",
			Callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back to previous map",
			Callback:    CommandMapb,
		},
	}
}
