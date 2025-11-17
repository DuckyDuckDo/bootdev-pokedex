package main

import (
	"time"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
)

// Main function controls the logic of the REPL
func main() {
	// Initialize the client
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Initialize the configs
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	// Starts the REPL where all of the main logic is performed
	startRepl(cfg)
}
