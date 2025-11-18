package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
)

// establish the struct for cliCommand
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, argument string) error
}

// ALL CLI COMMANDS REPL CAN CALL

// callback for the exit command
func commandExit(cfg *Config, empty_argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// call back for the help command
func commandHelp(cfg *Config, empty_argument string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// Loops through registry and prints out possible commands for use
	for key, value := range registry {
		fmt.Printf("%s: %s \n", key, value.description)
	}
	return nil
}

// call back for commanding map to go forward
func commandMapf(cfg *Config, empty_argument string) error {
	// Gets locations either through cache or API call based on configs next URL
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Updates URLs of the configs based on JSON response
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Prints out the locations
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// callback for moving map backwards, same as commandmapf
func commandMapb(cfg *Config, empty_argument string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// call back for exploring a specific location
func commandExplore(cfg *Config, location string) error {
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		fmt.Println("Location Not Found")
		return err
	}

	// Parses the JSON response and returns all found pokemon
	for _, encounter := range exploreResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

// call back for catching a specific pokemon
func commandCatch(cfg *Config, pokemon string) error {
	pokemonInfo, err := cfg.pokeapiClient.ViewPokemon(pokemon)
	if err != nil {
		fmt.Println("Pokemon does not exist!")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if checkCatch(pokemonInfo) {
		fmt.Printf("%s was caught! \n", pokemon)
		cfg.pokedex.Add(pokemon, pokemonInfo)
	}

	return nil
}

func checkCatch(pokemonInfo pokeapi.RespPokemonInfo) bool {
	baseXP := pokemonInfo.BaseExperience
	return baseXP > 1
}
