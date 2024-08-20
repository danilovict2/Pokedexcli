package main

import (
	"math/rand"
	"errors"
	"fmt"
	"os"

	"github.com/danilovict2/Pokedexcli/internal/pokeapi"
)

func commandInspect(conf *config, params []string) error {
	if len(params) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, ok := conf.Pokedex[params[0]]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}

func commandCatch(conf *config, params []string) error {
	if len(params) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	
	fmt.Println("Throwing a Pokeball at " + params[0] + "...")
	pokemon, err := pokeapi.GetPokemonData(params[0])

	if chance := rand.Intn(pokemon.BaseExperience) + 1; chance > 50 {
		fmt.Println(params[0] + " escaped!")
		return nil
	}  

	if err != nil {
		return err
	}

	fmt.Println(pokemon.Name + " was caught!")
	conf.Pokedex[pokemon.Name] = pokemon

	return nil
}

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
