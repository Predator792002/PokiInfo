package main

import "fmt"

func CommandPokeIndex(cfg *config, args []string) error {
	fmt.Println("Your PokeIndex:")
	for pokemon := range cfg.pokeIndex {
		fmt.Println(" - ", pokemon)
	}
	return nil
}
