package commands

import (
	"fmt"
	"log"

	"github.com/Kapung/gopokedex/pokeapi"
)

func CommandMap() {
	apiClient := pokeapi.NewClient()
	r, err := apiClient.PrintLocations()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locations:")
	for _, location := range r.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
}
