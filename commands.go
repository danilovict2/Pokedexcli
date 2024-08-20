package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/danilovict2/Pokedexcli/internal/pokeapi"
)	

func commandExplore(conf *config, params []string) error {
	if len(params) == 0 {
		return errors.New("you must provide a location name")
	}
	
	fmt.Println("Exploring " + params[0] + "...")

	resp, err := pokeapi.GetLocationAreas(params[0], &conf.Cache)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println("-" + pokemon.Data.Name)
	}

	return nil
}

func commandMapf(conf *config, params []string) error {
	resp, err := pokeapi.GetLocations(conf.Next, &conf.Cache)
	if err != nil {
		return err
	}

	conf.Next = resp.Next
	conf.Previous = resp.Previous
	
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(conf *config, params []string) error {
	if conf.Previous == nil {
		return errors.New("you're on the first page")
	}

	resp, err := pokeapi.GetLocations(conf.Previous, &conf.Cache)
	if err != nil {
		return err
	}

	conf.Next = resp.Next
	conf.Previous = resp.Previous
	
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandHelp(conf *config, params []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	
	commands := getCommands()

	for command := range commands {
		fmt.Println(commands[command].name + ": " + commands[command].description)
	}

	return nil
}

func commandExit(conf *config, params []string) error {
	os.Exit(0)
	return nil
}
