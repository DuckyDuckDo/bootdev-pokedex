package main

import (
	"errors"
	"fmt"
	"os"
)

// establish the struct for cliCommand
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, location string) error
}

// ALL CLI COMMANDS REPL CAN CALL

// callback for the exit command
func commandExit(cfg *Config, location string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// call back for the help command
func commandHelp(cfg *Config, location string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// Loops through registry and prints out possible commands for use
	for key, value := range registry {
		fmt.Printf("%s: %s \n", key, value.description)
	}
	return nil
}

// call back for commanding map to go forward
func commandMapf(cfg *Config, location string) error {
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
func commandMapb(cfg *Config, location string) error {
	if cfg.prevLocationsURL == nil {
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
