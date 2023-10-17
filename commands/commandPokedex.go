package commands

import (
	"fmt"

	"github.com/Kapung/gopokedex/config"
)

func CommandPokedex(cfg *config.Config, args ...string) {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Caught {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return
}
