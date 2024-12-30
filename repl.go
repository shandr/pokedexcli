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

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List all available commands",
			callback:    commandHelp,
		},
	}
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		command := cleanInput(text)
		value, exist := commands[command]
		if exist {
			err := value.callback()
			if err != nil {
				fmt.Println("An error occurred:", err)
			}
			continue
		}
		fmt.Println("Unknown command")
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
	for _, value := range commands {
		fmt.Printf("  %s: %s\n", value.name, value.description)
	}
	return nil
}