package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	if cfg.prevPageURL != nil && cfg.nextPageURL == nil {
		return errors.New("you're on the last page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = resp.Next
	cfg.prevPageURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevPageURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = resp.Next
	cfg.prevPageURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
