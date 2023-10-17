package commands

import "github.com/Kapung/gopokedex/config"

type cliCommand struct {
	name        string
	description string
	Callback    func(*config.Config, ...string)
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
		"explore": {
			name:        "explore [route_name]",
			description: "Explore the given area",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Try to catch the given Pokemon and add it to Pokedex",
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the stats of the Pokemon you've caught",
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List of all the Pokemon you've caught so far",
			Callback:    CommandPokedex,
		},
	}
}
