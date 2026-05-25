package main

func commandExplore(cfg *config, location string) error {
	pokemonResp, err := cfg.pokeapiClient.GetPokemonList(location)
}
