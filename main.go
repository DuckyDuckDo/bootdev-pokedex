package main

import (
	"time"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokedex"
)

// Main function controls the logic of the REPL
func main() {
	// Initialize the client
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Initialize the configs
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex.NewPokedex(),
	}

	// Starts the REPL where all of the main logic is performed
	startRepl(cfg)
}
