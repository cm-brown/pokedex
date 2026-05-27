package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("You must provide a location name\nusage: explore <location>")
	}

	location := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemonList(location)
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
