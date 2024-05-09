package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	config := initConfig()
	commands := getCommands(config)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex >")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandname := words[0]
		command, exists := commands[commandname]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	Next     *string
	Previous *string
}

func initConfig() *Config {
	nextURL := "https://pokeapi.co/api/v2/location-area/"
	// Previous is nil initially if there's no "back" at the start
	return &Config{
		Next:     &nextURL,
		Previous: nil, // Explicitly setting as nil, indicating there's no previous page initially
	}
}

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "20 locations",
			callback: func() error {
				return commandMap(config)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "20 locations",
			callback: func() error {
				return commandMapb(config)
			},
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
