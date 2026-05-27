package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(location string) (RespShallowPokemon, error) {
	url := baseURL + "/location-area/" + location

	if cached, ok := c.cache.Get(url); ok {
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

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
