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
	// data := pokeapi.GetLocationAreas()
	// areas, err := pokeapi.GetLocationAreas()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var previousUrl string
	// if areas.Previous == nil {
	// 	previousUrl = "There is no page to go back to"

	// } else {

	// 	previousUrl = *areas.Previous
	// }
	// nextUrl := *areas.Next
	// fmt.Printf("the prevUrl:%v, nextUlr:%v, all areas:%v", previousUrl, nextUrl, areas.Results)
	for {
		fmt.Print("Podekex > ")
		reader.Scan()

		words := clearInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			startTime := time.Now()
			err := command.callback(cfg)

			endTime := time.Now()
			executionTIme := endTime.Sub(startTime)
			fmt.Printf("Function executed in: %v/n", executionTIme)
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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandhelp,
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
