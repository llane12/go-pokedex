package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeApiClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapiClient: pokeApiClient,
	}

	startRepl(&cfg)
}
