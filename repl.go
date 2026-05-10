package main

import (
	"fmt"
	"os"
	"strings"
)

// Takes input, removes whitespace and converts to lowercase
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
