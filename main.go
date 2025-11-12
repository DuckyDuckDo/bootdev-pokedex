package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
)

var registry map[string]cliCommand


type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	for key, value := range registry {
		fmt.Printf("%s: %s \n", key, value.description)
	}

	return nil
}


func main() {
	// Initializes a scanner for the REPL
	scanner := bufio.NewScanner(os.Stdin)

	// Initializes registry of possible REPL commands
	registry = map[string]cliCommand{
				"exit": {
					name: "exit", 
					description: "Exit the pokedex",
					callback: commandExit,
				},
				"help": {
					name: "help",
					description: "Displays a help message", 
					callback: commandHelp,
				},
			}

	// Infinite for loop scanning for user input and doing something with it
	// Only exits on Ctrl + C from the user
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		command = strings.ToLower(command)
		command = strings.TrimSpace(command)
		switch command {
		case "exit":
			registry["exit"].callback()
		case "help":
			registry["help"].callback()
		default:
			fmt.Println("Unknown command")
		}

	}
}



