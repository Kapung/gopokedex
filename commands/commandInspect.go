package commands

import (
	"fmt"

	"github.com/Kapung/gopokedex/config"
)

func CommandInspect(cfg *config.Config, args ...string) {
	if len(args) != 1 {
		fmt.Println("No pokemon given")
		return
	}
	pokemonName := args[0]

	pkmn, ok := cfg.Caught[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\n", pkmn.Name, pkmn.Height, pkmn.Weight)
	for _, stat := range pkmn.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pkmn.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return
}
