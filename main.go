package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kapung/gopokedex/commands"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		command := scanner.Text()
		command = strings.ToLower(command)
		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Wrong command")
			continue
		}
		cmd.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commands.CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commands.CommandExit,
		},
	}
}
