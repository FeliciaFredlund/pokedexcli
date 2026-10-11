package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func commandExit(config *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)

	return fmt.Errorf("error exiting the program")
}

func commandHelp(config *config, _ string) error {
	fmt.Print("\nAvailable commands:\n\n")

	for _, cmd := range config.cmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *config, _ string) error {
	locations, next, previous, err := config.pokeapiClient.GetLocations(config.mapNextUrl)
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

func commandMapBack(config *config, _ string) error {
	if config.mapPreviousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, next, previous, err := config.pokeapiClient.GetLocations(config.mapPreviousUrl)
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

func commandExplore(config *config, area string) error {
	areaUrl := config.areaBaseUrl + area
	pokemons, err := config.pokeapiClient.GetPokemons(areaUrl)
	if err != nil {
		return fmt.Errorf("getting the area's pokemon data failed")
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found pokemon:")
	for _, pokemon := range pokemons {
		fmt.Println("-", pokemon)
	}

	return nil
}
