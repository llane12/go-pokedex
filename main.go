package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeApiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeApiClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
