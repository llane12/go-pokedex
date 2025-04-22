package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatchPokemon(cfg *config) error {
	if len(cfg.arguments) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokemonName := cfg.arguments[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.IntN(pokemon.BaseExperience) <= 40 {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
