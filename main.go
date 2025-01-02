package main

import (
	"github.com/shandr/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 60*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
