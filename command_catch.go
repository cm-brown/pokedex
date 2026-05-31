package main

import (
	"fmt"

	"github.com/cm-brown/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	pokemon := args[0]
	isValid := pokeapi.ValidPokemon(pokemon)
	if isValid == false {
		fmt.Printf("Invalid Pokemon: %s\nPlease enter a valid pokemon.\n", pokemon)
	} else {
		fmt.Printf("Throwing a pokeball at %s...\n", pokemon)
	}
	return nil
}
