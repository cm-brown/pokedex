package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex >")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		fmt.Printf("Your command was: %s\n", cleanLine[0])
	}
}
