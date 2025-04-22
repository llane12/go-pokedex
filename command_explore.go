package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.arguments) != 1 {
		return errors.New("you must provide a location name")
	}

	location, err := cfg.pokeapiClient.GetLocationArea(&cfg.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
