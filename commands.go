package main

import (
	"fmt"
	"os"
)

// The structure of command - it must have "name", "description" and "callback" function
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// os.Exit with code zero (0) -- indicates success
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getSupportedCommands() {
		fmt.Printf("	%s - %s\n", cmd.name, cmd.description)
	}
	return nil
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
