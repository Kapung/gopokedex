package commands

import (
	"fmt"
	"os"
)

func CommandExit() {
	fmt.Println("Exiting the Pokedex")
	os.Exit(0)
}
