package main

import (
	"errors"
	"fmt"
	"github.com/shandr/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]

	_, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	DisplayPokemonInfo(cfg.caughtPokemon[name])
	return nil
}

func DisplayPokemonInfo(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}
}
