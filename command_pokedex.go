package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}
	for name, _ := range cfg.caughtPokemon {
		fmt.Println("-", name)
	}
	return nil
}
