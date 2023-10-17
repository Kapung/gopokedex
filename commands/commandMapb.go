package commands

import (
	"fmt"
	"log"

	"github.com/Kapung/gopokedex/config"
)

func CommandMapb(cfg *config.Config) {

	if cfg.PreviousURL == nil {
		fmt.Println("You're already on the first page")
		return
	}

	r, err := cfg.PokeAPI.PrintLocations(cfg.PreviousURL)

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
