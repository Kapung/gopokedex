package config

import "github.com/Kapung/gopokedex/pokeapi"

type Config struct {
	PokeAPI     pokeapi.Client
	NextURL     *string
	PreviousURL *string
}
