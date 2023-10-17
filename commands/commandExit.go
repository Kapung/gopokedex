package commands

import (
	"fmt"
	"os"

	"github.com/Kapung/gopokedex/config"
)

func CommandExit(cfg *config.Config) {
	fmt.Println("Exiting the Pokedex")
	os.Exit(0)
}
