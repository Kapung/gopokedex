package main

import (
	"time"

	"github.com/Kapung/gopokedex/config"
	"github.com/Kapung/gopokedex/pokeapi"
)

func main() {
	cfg := &config.Config{
		PokeAPI: pokeapi.NewClient(time.Minute * 10),
	}
	startProgram(cfg)
}
