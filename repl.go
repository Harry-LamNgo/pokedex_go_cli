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
	"strings"
)

func cleanInput(text string) []string {
	var words []string
	words = strings.Fields(strings.ToLower(text))
	return words
	// return []string{}
}
