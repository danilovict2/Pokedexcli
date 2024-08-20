package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/danilovict2/Pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func GetLocations(pageURL *string, c *pokecache.Cache) (RespShallowLocations, error) {
	url := BaseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)

		if err != nil {
			return RespShallowLocations{}, err
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return RespShallowLocations{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
		}

		if err != nil {
			return RespShallowLocations{}, err
		}

		data = body
		c.Add(url, data)
	}

	locations := RespShallowLocations{}
	json.Unmarshal(data, &locations)

	return locations, nil
}

func GetLocationAreas(name string, c *pokecache.Cache) (RespLocationAreas, error) {
	url := BaseURL + fmt.Sprintf("location-area/%s", name)
	
	data, ok := c.Get(url)

	if !ok {
		res, err := http.Get(url)

		if err != nil {
			return RespLocationAreas{}, err
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return RespLocationAreas{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
		}

		if err != nil {
			return RespLocationAreas{}, err
		}

		data = body
		c.Add(url, data)
	}
	
	locationAreas := RespLocationAreas{}
	json.Unmarshal(data, &locationAreas)

	return locationAreas, nil
}
