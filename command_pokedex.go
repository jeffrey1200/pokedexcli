package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	// var pokemons string
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("you haven't caught any pokemons yet")
		return nil
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("-", pokemon.Name)
	}

	return nil
}
