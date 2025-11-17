package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations - GET REQUEST to pokeapi
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// Build out URL
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Try hitting the cache first
	if data, ok := c.pokedexCache.Get(url); ok {
		var locationsResp RespShallowLocations
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	// Go to HTTP if data not found in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// Cache and unmarshal the response
	dat, err := io.ReadAll(resp.Body)
	c.pokedexCache.Add(url, dat)

	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
