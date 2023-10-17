package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kapung/gopokedex/commands"
	"github.com/Kapung/gopokedex/pokeapi"
)

type config struct {
	pokeAPI     pokeapi.Client
	nextURL     *string
	previousURL *string
}

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	cmdMap := commands.GetCommands()
	cfg := config{
		pokeAPI: pokeapi.NewClient(),
	}

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
		cmd.Callback()
	}
}
