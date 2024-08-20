package pokeapi

type RespShallowLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

type RespLocationAreas struct {
	PokemonEncounters []struct {
		Data struct {
			Name string `json:"name"`
 			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}