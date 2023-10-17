package commands

import "fmt"

func CommandHelp() {
	commands := GetCommands()

	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println("")
}
