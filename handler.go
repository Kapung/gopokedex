package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kapung/gopokedex/commands"
	"github.com/Kapung/gopokedex/config"
)

func startProgram(cfg *config.Config) {
	scanner := bufio.NewScanner((os.Stdin))
	cmdMap := commands.GetCommands()

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		command := scanner.Text()
		command = strings.ToLower(command)
		cmd, ok := cmdMap[command]
		if !ok {
			fmt.Println("Wrong command")
			continue
		}
		cmd.Callback(cfg)
	}
}
