package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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

		// Process input
		cleanInput := cleanInput(inputText)

		fmt.Printf("Your command was: %v \n", cleanInput[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("reading standard input: %v", err)
	}
}
