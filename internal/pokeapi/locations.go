package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/danilovict2/Pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func GetLocations(pageURL *string, c *pokecache.Cache) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
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
