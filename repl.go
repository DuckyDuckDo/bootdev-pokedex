package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
)

// declaring global variable registry to map commands to cliCommand interface
var registry map[string]cliCommand

// establish the struct for Config
type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *Config) {
	// Initialize a scanner for the REPL
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize registry of possible REPL commands
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Dislays locations using Poke-API",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display locations but goes backward",
			callback:    commandMapb,
		},
	}

	// Infinite for loop scanning for user input and doing something with it
	// Only exits on Ctrl + C from the user
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// cleans user input
		command := scanner.Text()
		command = strings.ToLower(command)
		command = strings.TrimSpace(command)

		// perform command based on the registry of available commands
		switch command {
		case "exit":
			registry["exit"].callback(cfg)
		case "help":
			registry["help"].callback(cfg)
		case "map":
			registry["map"].callback(cfg)
		case "mapb":
			registry["mapb"].callback(cfg)
		default:
			fmt.Println("Unknown command")
		}
	}
}
