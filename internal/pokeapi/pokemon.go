package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokemonData(name string) (Pokemon, error) {
	url := BaseURL + "pokemon/" + name
	res, err := http.Get(url)

	if err != nil {
		return Pokemon{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	json.Unmarshal(body, &pokemon)

	return pokemon, nil
}
