package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
	pokemonEncounters, err := config.pokeapiClient.GetPokemonsEncounters(areaUrl)
	if err != nil {
		return fmt.Errorf("getting the area's pokemon data failed")
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found pokemon:")
	for _, pokemon := range pokemonEncounters {
		fmt.Println("-", pokemon)
	}

	return nil
}

func commandCatch(config *config, pokemonName string) error {
	pokemonUrl := config.pokemonBaseUrl + pokemonName
	pokemon, err := config.pokeapiClient.GetPokemon(pokemonUrl)
	if err != nil || pokemon.Name == "" {
		return fmt.Errorf("getting the pokemon's data failed")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Randomly deciding if pokemon was caught.
	// Number between 0 and pokemon's base experience is generated
	// If it is below or equal to 80 it is caught
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := r.Intn(pokemon.BaseExperience)
	if number > 80 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	config.pokedex[pokemonName] = pokemon

	return nil
}
