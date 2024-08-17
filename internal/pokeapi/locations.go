package pokeapi

import (
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

func GetLocations(pageURL *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area";
	if pageURL != nil {
		url = *pageURL
	}

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

	locations := RespShallowLocations{}
	json.Unmarshal(body, &locations)

	return locations, nil
}