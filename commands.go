package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)

	return fmt.Errorf("error exiting the program")
}

func commandHelp(config *config) error {
	fmt.Print("\nAvailable commands:\n\n")

	for _, cmd := range config.cmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *config) error {
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

func commandMapBack(config *config) error {
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
