package main

import (
	"time"

	"github.com/patrickneise/pokedexcli/internal/pokeapi"
	"github.com/patrickneise/pokedexcli/internal/pokerepl"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := pokerepl.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	pokerepl.StartREPL(&cfg)
}
