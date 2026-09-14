package main

import (
	"fmt"
	"os"
)

// The structure of command - it must have "name", "description" and "callback" function
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Create structure of PokeAPI-location-area
type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Method wrap commands in a function --> this function only run after main() run
//   --> re-create map of commands each time this function call -> cannot accidentally delete the command like declare only one time by global var

func getSupportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// os.Exit with code zero (0) -- indicates success
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.commands {
		fmt.Printf("	%s - %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"

	if cfg.nextURL != nil {
		locationAreaURL = *cfg.nextURL
	}

	fetchLocationArea(locationAreaURL, cfg)

	return nil
}

func commandMapb(cfg *config) error {

	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchLocationArea(*cfg.previousURL, cfg)
}

// Method declare Global var -- Go auto runs init() before main() -> add all commands into one map and shared everywhere in package main
//   --> this map lives for entire duration of the program
/*
var supportedCommands = map[string]cliCommand{}

func init() {
	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}
*/
