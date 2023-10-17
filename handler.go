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
		input := scanner.Text()
		parts := strings.Split(input, " ")
		command := strings.ToLower(parts[0])
		args := parts[1]

		cmd, ok := cmdMap[command]

		if !ok {
			fmt.Println("Wrong command")
			continue
		}
		cmd.Callback(cfg, args)
	}
}
