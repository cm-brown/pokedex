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
		if cleanLine[0] == "exit" {
			commandExit()
		}
	}
}
