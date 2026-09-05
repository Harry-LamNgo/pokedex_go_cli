package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Combine both bufio.NewScanner and os.Stdin
	// --> whenever you call scanner.Scan, it will block and wait for User's input

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		var inputText string
		for scanner.Scan() {
			inputText = scanner.Text()
			cleanInput := cleanInput(inputText)
			fmt.Printf("Your command was: %v \n", cleanInput[0])
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("reading standard input: %v", err)
		}

	}
}
