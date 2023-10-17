package commands

import (
	"fmt"
	"math/rand"

	"github.com/Kapung/gopokedex/config"
)

func CommandCatch(cfg *config.Config, args ...string) {
	if len(args) != 1 {
		fmt.Println("No pokemon given")
		return
	}
	pokemonName := args[0]

	if _, alreadyCaught := cfg.Caught[pokemonName]; alreadyCaught {
		fmt.Printf("%s was already caught!\n", pokemonName)
		return
	}

	pkmn, err := cfg.PokeAPI.GetPokemon(pokemonName)

	if err != nil {
		fmt.Println("There is no Pokemon by that name")
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pkmn.Name)

	const success = 50
	catchRate := rand.Intn(pkmn.BaseExperience)

	if catchRate > success {
		fmt.Printf("%s escaped!\n", pkmn.Name)
		return
	}

	cfg.Caught[pkmn.Name] = pkmn
	fmt.Printf("%s was caught!\n", pkmn.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	return
}
