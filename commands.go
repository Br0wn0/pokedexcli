package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Br0wn0/pokedexcli/internal/pokeapi"
)

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Use the following commands: 'exit', 'map', 'mapb'")
	return nil
}

func commandMap(config *Config) error {
	if config.Next == nil {
		return errors.New("no next url available")
	}
	data, err := pokeapi.ProcessData(*config.Next)
	if err != nil {
		log.Printf("failed to process data: %v", err)
		return err
	}
	for _, result := range data.Results {
		fmt.Printf("Name: %s, Url: %s\n", result.Name, result.URL)
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

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("no previous url available")
	}
	data, err := pokeapi.ProcessData(*config.Previous)
	if err != nil {
		log.Printf("failed to process data: %v", err)
		return err
	}
	for _, result := range data.Results {
		fmt.Printf("Name: %s, Url: %s\n", result.Name, result.URL)
	}
	if data.Previous != nil {
		config.Previous = data.Previous
	}
	config.Next = data.Next
	return nil
}
