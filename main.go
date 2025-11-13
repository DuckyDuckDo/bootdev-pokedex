package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// declaring global variable registry to map commands to cliCommand interface
var registry map[string]cliCommand

// establish the struct for cliCommand
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

// establish the struct for Config
type Config struct {
	Next     string
	Previous string
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// callback for the exit command
func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// call back for the help command
func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// Loops through registry and prints out possible commands for use
	for key, value := range registry {
		fmt.Printf("%s: %s \n", key, value.description)
	}

	return nil
}

func commandMap(cfg *Config) error {
	var url string
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = cfg.Next
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	locations := Locations{}
	if err = json.Unmarshal(data, &locations); err != nil {
		log.Fatal(err)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, location_struct := range locations.Results {
		fmt.Printf("%v \n", location_struct.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	var url string
	if cfg.Previous == "" {
		fmt.Println("You're on the first page.")
		return nil
	} else {
		url = cfg.Previous
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	locations := Locations{}
	if err = json.Unmarshal(data, &locations); err != nil {
		log.Fatal(err)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, location_struct := range locations.Results {
		fmt.Printf("%v \n", location_struct.Name)
	}
	return nil

}

func main() {
	// Initializes a scanner for the REPL
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize a Config Struct that will maintain next and previous urls
	cfg := &Config{
		Next:     "",
		Previous: "",
	}

	// Initializes registry of possible REPL commands
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
			callback:    commandMap,
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
