package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Start REPL
	for {
		fmt.Println("Pokedex >")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)

		// Check if user command exists
		cmd, ok := getCommands()[cleanLine[0]]
		if !ok {
			fmt.Println("Unknown command:", cleanLine[0])
			continue
		}
		// Executes command callback
		err := cmd.callback()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
