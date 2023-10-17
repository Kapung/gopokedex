package commands

import (
	"fmt"
	"log"

	"github.com/Kapung/gopokedex/config"
)

func CommandExplore(cfg *config.Config, args ...string) {
	if len(args) != 1 {
		fmt.Println("No route given")
		return
	}
	routeName := args[0]

	r, err := cfg.PokeAPI.PrintRoute(routeName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exploring %s...\n", r.Name)
	fmt.Println("Found Pokemon:")
	for _, route := range r.PokemonEncounters {
		fmt.Printf(" - %s\n", route.Pokemon.Name)
	}
	return
}
