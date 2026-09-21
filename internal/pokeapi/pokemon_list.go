package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemons(areaName string) (PokemonAreaResponse, error) {
	url := baseURL + "/location-area/" + areaName

	// Check cacheData exist - if yes skip request and use cacheData to Unmarshal
	cacheData, exist := c.cache.Get(url)
	if !exist {
		// New Request - GET Method - send request to URL
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonAreaResponse{}, err
		}

		// Do - Get the response from URL
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonAreaResponse{}, fmt.Errorf("Failed to fetch the Pokemons in area (%v)", areaName)
		}
		defer resp.Body.Close()

		// Read the body of fetched data
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return PokemonAreaResponse{}, fmt.Errorf("Failed to read the data: %w", err)
		}

		// Add data and url to future cache
		c.cache.Add(url, data)
		cacheData = data
	}

	// Unmarshal the data -- the data in this case is single struct -> NO [] or NO slice
	var dataResponse PokemonAreaResponse
	if err := json.Unmarshal(cacheData, &dataResponse); err != nil {
		return PokemonAreaResponse{}, err
	}

	if len(dataResponse.PokemonEncounters) == 0 {
		return PokemonAreaResponse{}, fmt.Errorf("No Pokemon found in %s \n", areaName)
	}

	return dataResponse, nil
}
