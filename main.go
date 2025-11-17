package main

import (
	"time"
	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
)



// Main function controls the logic of the REPL
func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
