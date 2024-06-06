package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	isPokemonCatched := tryToCatch(pokemon.BaseExperience)
	fmt.Printf("Throwing a pokeball at %s...", pokemon.Name)
	if !isPokemonCatched {

		fmt.Printf("%s has escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s has been caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func calculateCatchRate(baseExperience, minExperience, maxExperience int) float64 {
	normalizeExperience := float64(baseExperience-minExperience) / float64(maxExperience-minExperience)
	catchRate := 1 - normalizeExperience
	return catchRate
}

func tryToCatch(baseExperience int) bool {
	minExperience := 50
	maxExperience := 340
	catchRate := calculateCatchRate(baseExperience, minExperience, maxExperience)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomThreshold := rng.Float64()

	fmt.Printf("random threshold:%0.3f, catch rate: %f", randomThreshold, catchRate)
	if randomThreshold < catchRate {
		return true
	} else {
		return false
	}

}
