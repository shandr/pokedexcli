package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func repl() {
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
			err := value.callback()
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
	output := strings.TrimSpace(text)
	output = strings.ToLower(text)
	words := strings.Fields(output)
	commands := strings.Join(words, " ")
	return commands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Available commands:")
	for _, value := range getCommands() {
		fmt.Printf("  %s: %s\n", value.name, value.description)
	}
	return nil
}
