package main

import (
	"time"

	"github.com/jeffrey1200/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.PokemonInformation{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
