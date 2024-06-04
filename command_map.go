package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationResponse, err := cfg.pokeapiClient.ListLocation(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResponse.Next
	cfg.prevLocationsURL = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResponse, err := cfg.pokeapiClient.ListLocation(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResponse.Next
	cfg.prevLocationsURL = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
