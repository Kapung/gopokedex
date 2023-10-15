package commands

import "fmt"

func CommandHelp() {
	commands := getCommands()

	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println("")
}

type cliCommand struct {
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
		},
	}
}
