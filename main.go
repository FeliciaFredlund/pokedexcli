package main

import (
	"time"

	"github.com/FeliciaFredlund/pokedexcli/internals/pokeapi"
)

type config struct {
	cmds           map[string]cliCommand
	pokeapiClient  pokeapi.Client
	mapNextUrl     string
	mapPreviousUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	config := &config{
		cmds:          getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
		mapNextUrl:    "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	replLoop(config)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "Help",
			description: "Displays available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "Displays the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "MapB",
			description: "Displays the previous 20 areas",
			callback:    commandMapBack,
		},
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
