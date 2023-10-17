package commands

import (
	"fmt"
	"os"

	"github.com/Kapung/gopokedex/config"
)

func CommandExit(cfg *config.Config, args ...string) {
	fmt.Println("Exiting the Pokedex")
	os.Exit(0)
}
