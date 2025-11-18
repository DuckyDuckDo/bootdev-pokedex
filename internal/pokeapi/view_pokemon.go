package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ViewPokemon(pokemon string) (RespPokemonInfo, error) {
	// Build URL
	url := pokemonURL + "/" + pokemon

	// Try viewing from cache first
	if data, ok := c.pokeapiCache.Get(url); ok {
		var pokemonResp RespPokemonInfo
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return RespPokemonInfo{}, err
		}
		return pokemonResp, nil
	}

	// Make the HTTP Get Request Call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer resp.Body.Close()

	// Cache and unmarshal the response into JSON Data
	dat, err := io.ReadAll(resp.Body)
	c.pokeapiCache.Add(url, dat)

	if err != nil {
		return RespPokemonInfo{}, err
	}

	var pokemonResp RespPokemonInfo
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	return pokemonResp, nil

}
