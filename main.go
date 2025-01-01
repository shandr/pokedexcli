package main

import (
	"github.com/shandr/pokedexcli/internal/pokeapi"
)

func main() {
	config := &pokeapi.Config{}
	repl(config)
}
