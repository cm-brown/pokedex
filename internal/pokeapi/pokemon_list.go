package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/cm-brown/pokedex/internal/pokecache"
)

func (c *Client) GetPokemonList(location string, pageURL *string, cache *pokecache.Cache) (RespShallowPokemon, error) {
	url := baseURL + "/location-area/" + location
	if pageURL != nil {
		url = *pageURL
	}

	if cached, ok := cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(cached, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	pokemonResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	return pokemonResp, nil
}
