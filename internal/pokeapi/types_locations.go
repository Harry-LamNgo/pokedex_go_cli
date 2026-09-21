package pokeapi

// Create type structure of PokeAPI-location-area
type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Create type structure of PokeAPI Pokemon Encounter in Area - Including what you must use
type PokemonAreaResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
