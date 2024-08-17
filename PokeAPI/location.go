package pokeapi

import (
	"io"
	"net/http"
	"fmt"
	"errors"
)

func GetLocation(id int) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + string(id))

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		errors.New(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
	}
	if err != nil {
		return err
	}

	fmt.Printf("%s", body)
	return nil
}