package main

import (
	"fmt"

	"github.com/cm-brown/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args []string) error {
	_, pokemon := pokeapi.ValidPokemon(args[0])
	if _, ok := cfg.pokedex[args[0]]; ok {
		{
			fmt.Printf("Name: %s\n", pokemon.Name)
			fmt.Printf("Height: %d\n", pokemon.Height)
			fmt.Println("Status:")
			for _, stats := range pokemon.Stats {
				fmt.Printf("- %s: %d\n", stats.Stat.Name, stats.BaseStat)
			}
			fmt.Println("Types:")
			for _, types := range pokemon.Types {
				fmt.Printf("- %s\n", types.Type.Name)
			}
		}

	}
	return nil
}
