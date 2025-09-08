package main

import (
	"time"

	"github.com/Predator792002/PokiInfo/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Minute, 5*time.Minute)

	cfg := &config{
		pokeIndex:     map[string]pokeapi.CaughtPokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
