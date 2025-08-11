package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FeliciaFredlund/pokedexcli/internals/pokeapi"
)

type config struct {
	mapNextUrl     string
	mapPreviousUrl string
}

func replLoop() {
	reader := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	config := &config{
		mapNextUrl: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		if command, exists := cmds[input[0]]; exists {
			if err := command.callback(config); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)

	return fmt.Errorf("error exiting the program")
}

func commandHelp(config *config) error {
	fmt.Print("\nAvailable commands:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *config) error {
	locations, next, previous, err := pokeapi.GetLocations(config.mapNextUrl)
	if err != nil {
		return fmt.Errorf("getting location data failed")
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	config.mapNextUrl = next
	config.mapPreviousUrl = previous

	return nil
}

func commandMapBack(config *config) error {
	if config.mapPreviousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, next, previous, err := pokeapi.GetLocations(config.mapPreviousUrl)
	if err != nil {
		return fmt.Errorf("getting location data failed")
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	config.mapNextUrl = next
	config.mapPreviousUrl = previous

	return nil
}
