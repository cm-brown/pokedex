package main

import (
	"fmt"
	"os"
	"strings"
)

// Command structure
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandList = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exits the Pokedex",
		callback:    commandExit,
	},
}

// Takes input, removes whitespace and converts to lowercase
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// Exits program
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
