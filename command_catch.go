package main

import (
	"fmt"
	"math/rand"

	"github.com/cm-brown/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	pokemon := args[0]
	isValid, pokeresp := pokeapi.ValidPokemon(pokemon)
	if isValid {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
		randInt := rand.Intn(50)
		result := randInt >= int(float64(pokeresp.BaseExperience)*0.1)
		if result {
			fmt.Printf("%s was caught!\n", pokemon)
			cfg.pokedex[pokeresp.Name] = pokeresp
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	} else {
		fmt.Printf("Invalid Pokemon: %s\nPlease enter a valid pokemon.\n", pokemon)
	}
	return nil
}
