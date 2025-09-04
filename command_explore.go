package main

import (
	"fmt"
	"strings"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		fmt.Println("usage: explore <area_name>")
		return nil
	}

	area := strings.ToLower(strings.TrimSpace(args[0]))

	fmt.Printf("Exploring %s...\n", area)
	pokemons, err := cfg.pokeapiClient.GetPokemons(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, name := range pokemons {
		fmt.Printf(" - %s\n", strings.ToLower(strings.TrimSpace(name)))
	}
	return nil
}
