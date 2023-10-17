package commands

import (
	"fmt"

	"github.com/Kapung/gopokedex/config"
)

func CommandHelp(cfg *config.Config, args ...string) {
	commands := GetCommands()

	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println("")
}
