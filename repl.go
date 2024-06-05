package main

import (
	"bufio"
	"fmt"
	"time"

	// "log"
	"os"
	"strings"

	"github.com/jeffrey1200/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Podekex > ")
		reader.Scan()

		words := clearInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]

		if exists {

			startTime := time.Now()
			err := command.callback(cfg, args...)

			endTime := time.Now()
			// test, er := cfg.pokeapiClient.GetPokemonsInArea("sunyshore-city-area")
			// fmt.Printf("these are the pokemons: %v, error if any:%v", test.Pokemon_encounters, er)
			executionTIme := endTime.Sub(startTime)
			fmt.Printf("Function executed in: %v/n, second command: %v", executionTIme, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func clearInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandhelp,
		},
		"explore": {
			name:        "explore",
			description: "Display pokemons in the given location",
			callback:    commandExplore,
		},

		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
