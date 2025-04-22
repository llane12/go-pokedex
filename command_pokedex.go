package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon yet. Get out there and start catching Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}
