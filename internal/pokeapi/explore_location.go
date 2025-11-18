package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Build out a GET request to explore a specific location
func (c *Client) ExploreLocation(location string) (RespExploration, error) {
	// Build URL for exploration
	url := exploreURL + "/" + location

	// Try hitting the cache first
	if data, ok := c.pokeapiCache.Get(url); ok {
		var exploreResp RespExploration
		err := json.Unmarshal(data, &exploreResp)
		if err != nil {
			return RespExploration{}, err
		}
		return exploreResp, nil
	}

	// Make the HTTP Get Request Call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploration{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploration{}, err
	}
	defer resp.Body.Close()

	// Cache and unmarshal the response into JSON Data
	dat, err := io.ReadAll(resp.Body)
	c.pokeapiCache.Add(url, dat)

	if err != nil {
		return RespExploration{}, err
	}

	var exploreResp RespExploration
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespExploration{}, err
	}

	return exploreResp, nil

}
