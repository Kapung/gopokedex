package commands

import (
	"fmt"
	"log"

	"github.com/Kapung/gopokedex/config"
)

func CommandMap(cfg *config.Config, args ...string) {
	r, err := cfg.PokeAPI.PrintLocations(cfg.NextURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locations:")
	for _, location := range r.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.NextURL = r.Next
	cfg.PreviousURL = r.Previous

	return
}
