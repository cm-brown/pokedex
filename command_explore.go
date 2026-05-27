package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	location := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemonList(location, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
