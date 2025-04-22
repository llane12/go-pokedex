package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextPageURL   *string
	prevPageURL   *string
	arguments     []string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the next 20 location areas",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Lists the Pokemon that can be enocuntered in a location area",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommands() {
		fmt.Println(value.name, ": ", value.description)
	}
	fmt.Println()
	return nil
}
