package main

import (
	"errors"
	"fmt"
	"math/rand"
	//"slices"
	"time"
)

//var caughtPokemons = []string{}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	//if slices.Contains(caughtPokemons, name) {
	//	fmt.Println("This pokemon already caught")
	//	return nil
	//}
	_, exists := cfg.caughtPokemon[name]
	if exists {
		fmt.Println("This pokemon already caught")
		return nil
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonExp := pokemon.BaseExperience
	catched := catchPokemon(pokemonExp)
	if catched {
		fmt.Printf("%s was caught!\n", name)
		//caughtPokemons = append(caughtPokemons, name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", name)
	return nil
}

// catchPokemon attempts to catch a Pokémon based on its base experience.
func catchPokemon(baseExperience int) bool {
	// Create a new random source and generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Calculate catch probability
	catchProbability := 1.0 - (float64(baseExperience) / 300.0)

	// Ensure minimum 50% catch rate
	if catchProbability < 0.5 {
		catchProbability = 0.5
	}

	// Generate a random number between 0 and 1
	randomValue := rng.Float64()

	// Pokémon is caught if the random value is less than the catch probability
	return randomValue < catchProbability
}
