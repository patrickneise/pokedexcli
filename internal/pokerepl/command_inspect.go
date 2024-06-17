package pokerepl

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.CaughtPokemon[name]
	if !ok {
		fmt.Printf("you have not caught %s", pokemon.Name)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}
	fmt.Println()

	return nil
}
