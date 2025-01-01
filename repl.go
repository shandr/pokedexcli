package main

import (
	"bufio"
	"fmt"
	"github.com/shandr/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the location areas forward",
			callback:    getMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the location areas backward",
			callback:    getMapBackward,
		},
	}
}

func repl(config *pokeapi.Config) {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		command := cleanInput(text)
		value, exist := getCommands()[command]
		if exist {
			err := value.callback(config)
			if err != nil {
				fmt.Println("An error occurred:", err)
			}
			continue
		}
		fmt.Println("Unknown command")
		continue
	}
}

func cleanInput(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return strings.Join(words, " ")
}

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config) error {
	fmt.Println("Available commands:")
	for _, value := range getCommands() {
		fmt.Printf("  %s: %s\n", value.name, value.description)
	}
	return nil
}

func getMap(config *pokeapi.Config) error {
	locations, err := pokeapi.GetLocationAreas(config)
	if err != nil {
		fmt.Println("Error getting location areas", err)
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name, location.URL)
	}
	return nil
}

func getMapBackward(config *pokeapi.Config) error {
	locations, err := pokeapi.GetLocationAreasBackward(config)
	if err != nil {
		fmt.Println("Error getting location areas", err)
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name, location.URL)
	}
	return nil
}
