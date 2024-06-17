package pokerepl

import (
	"fmt"
)

func CommandPokedex(cfg *Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s", pokemon.Name)
	}
	fmt.Println()

	return nil
}
