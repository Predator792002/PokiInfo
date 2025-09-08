package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func CommandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}
	name := strings.ToLower(strings.TrimSpace(args[0]))

	if _, ok := cfg.pokeIndex[name]; ok {
		fmt.Printf("%s is already in your Pokedex\n", name)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	p, err := cfg.pokeapiClient.CatchPokemons(name)
	if err != nil {
		fmt.Println("couldn't find that pokemon")
		return nil
	}
	prob := 1.0 - float64(p.BaseExperience)/400.
	if prob < 0.1 {
		prob = 0.1
	} else if prob > 0.9 {
		prob = 0.9
	}

	if rand.Float64() < prob {
		cfg.pokeIndex[p.Name] = p
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
