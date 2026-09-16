/*
The name of this file is 'REPL' stand for
- Read-Eval-Print-Loop -
Reads input from the user (like command or expression)
Evaluates that input (runs it, processes it, etc)
Prints the result back to the user
Loops back to step 1 and waits for next input
*/

// This function split the Use's text into words based on whitespace
// Also lowercase the input and trim any leading or trailing whitespace
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Harry-LamNgo/pokdex_go_cli/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

func cleanInput(text string) []string {
	var words []string
	words = strings.Fields(strings.ToLower(text))
	return words
}

func runREPL(cfg *config) {
	// Combine both bufio.NewScanner and os.Stdin--> whenever you call scanner.Scan, it will block and wait for User's input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Need to show "Pokedex > " prompt everytime a command is executed -> wrap that in manual for loop
		fmt.Print("Pokedex > ")

		var inputText string
		if scanner.Scan() == false {
			break
		}
		inputText = scanner.Text()

		// Process input -> and take 1st word
		cleanInput := cleanInput(inputText)[0]

		cmd, found := cfg.commands[cleanInput]
		if found == false {
			fmt.Println("Unknown command")
			continue
		} else {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("reading standard input: %v", err)
	}

}
