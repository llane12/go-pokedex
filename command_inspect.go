package main

import (
	"errors"
	"fmt"
	"sort"
)

func commandInspect(cfg *config) error {
	if len(cfg.arguments) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokemonName := cfg.arguments[0]

	pokemon, found := cfg.pokedex[pokemonName]
	if !found {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	sortedTypes := pokemon.Types[:]
	sort.Slice(sortedTypes, func(i, j int) bool {
		return sortedTypes[i].Slot < sortedTypes[j].Slot
	})
	for _, typ := range sortedTypes {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}
