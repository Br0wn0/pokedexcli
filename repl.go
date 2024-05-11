package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Br0wn0/pokedexcli/internal/pokeapi"
)

func repl(c *Config) {
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
			err := command.callback(c, words[1:]...)
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
	callback    func(config *Config, args ...string) error
}

type Config struct {
	Next     *string
	Previous *string
	PKD      map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	return &Config{
		PKD: make(map[string]pokeapi.Pokemon),
	}
}

func initConfig() *Config {
	nextURL := "https://pokeapi.co/api/v2/location-area/"
	return &Config{
		Next:     &nextURL,
		Previous: nil,
	}
}

func getlocationsurl(c *Config, locationName string) {
	baseAPIURL := "https://pokeapi.co/api/v2/location-area/"

	fullURL := baseAPIURL + locationName
	c.Next = &fullURL

}

func getpokemonurl(c *Config, pokemonName string) {
	baseAPIURL := "https://pokeapi.co/api/v2/pokemon/"

	fullURL := baseAPIURL + pokemonName
	c.Next = &fullURL
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "explore the pokemon in the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "inspect your pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "inspects pokedex",
			callback:    commandPokeDex,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
