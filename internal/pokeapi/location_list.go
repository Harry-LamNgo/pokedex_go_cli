package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Create method/feature of Client struct - name FetchLocations

func (c *Client) FetchLocations(pageURL *string) (LocationAreaResponse, error) {

	url := baseURL + "/location-area"

	// Check if pageURL (current URL is nil or not) -> if first time - no page -> use the default url
	if pageURL != nil {
		url = *pageURL
	}

	// Check cacheData before exist - if yes skip request and use the cacheDat to unmarshal

	cacheData, exist := c.cache.Get(url)
	if !exist {
		// New Request - GET Method - send request to URL
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		// Do - Get the response from URL
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Failed to fetch Pokemon Location-area: %w", err)
		}

		defer resp.Body.Close()

		// Read the body of fetched data
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Failed to read the data: %w", err)
		}

		// Add data and url to future cache
		c.cache.Add(url, data)
		cacheData = data
	}

	// Unmarshal the data -- the data in this case is a single struct
	var dataResponse LocationAreaResponse
	if err := json.Unmarshal(cacheData, &dataResponse); err != nil {
		return LocationAreaResponse{}, err
	}

	// Handle empty dataResponse.Results
	if len(dataResponse.Results) == 0 {
		fmt.Println("No location areas found")
		return LocationAreaResponse{}, nil
	}

	return dataResponse, nil
}
