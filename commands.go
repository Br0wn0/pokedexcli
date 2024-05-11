package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/Br0wn0/pokedexcli/internal/pokeapi"
)

func commandExit(config *Config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args ...string) error {
	fmt.Println("Use the following commands: 'exit', 'map', 'mapb'")
	return nil
}

func commandMap(config *Config, args ...string) error {
	if config.Next == nil {
		return errors.New("no next url available")
	}
	data, err := pokeapi.ProcessData(*config.Next)
	if err != nil {
		log.Printf("failed to process data: %v", err)
		return err
	}
	for _, result := range data.Results {
		fmt.Printf("Name: %s\n", result.Name)
	}
	if data.Next != nil {
		config.Next = data.Next
	} else {
		config.Next = nil
	}
	if data.Previous != nil {
		config.Previous = data.Previous
	}
	return nil

}

func commandMapb(config *Config, args ...string) error {
	if config.Previous == nil {
		return errors.New("no previous url available")
	}
	data, err := pokeapi.ProcessData(*config.Previous)
	if err != nil {
		log.Printf("failed to process data: %v", err)
		return err
	}
	for _, result := range data.Results {
		fmt.Printf("Name: %s\n", result.Name)
	}
	if data.Previous != nil {
		config.Previous = data.Previous
	}
	config.Next = data.Next
	return nil
}

func commandExplore(c *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("location name is required")
	}
	locationName := args[0]
	getlocationsurl(c, locationName)
	data, err := pokeapi.ProcessLocationData(*c.Next)
	if err != nil {
		log.Printf("failed to process data: %v", err)
		return err
	}
	for _, result := range data.PokemonEncounters {
		fmt.Printf("Found Pokemon: %s\n", result.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("?????")
	}
	pokemonName := args[0]
	getpokemonurl(c, pokemonName)
	data, err := pokeapi.CatchEmAll(*c.Next)
	if err != nil {
		return err
	}
	res := rand.Intn(data.BaseExperience)

	fmt.Printf("throwing a pokeball at %s...\n", data.Name)
	if res > 40 {
		fmt.Printf("IT ESCAPED, get THAT MDFKr")
		return nil
	}
	fmt.Printf("%s get caught you HAHAHAHA\n", data.Name)

	if c.PKD == nil {
		c.PKD = make(map[string]pokeapi.Pokemon)
	}

	c.PKD[data.Name] = data

	return nil

}

func commandInspect(c *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("INSPECT WHAT????")
	}
	pokemonName := args[0]
	pokemon, ok := c.PKD[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}

func commandPokeDex(c *Config, args ...string) error {
	if c.PKD == nil {
		fmt.Println("NO POKEDEX BLOKE")
	} else if len(c.PKD) == 0 {
		fmt.Println("Catch a pokemon LMAO u poor scrub")
	} else {
		for pokemonName := range c.PKD {
			fmt.Println(pokemonName)
		}
	}
	return nil
}
