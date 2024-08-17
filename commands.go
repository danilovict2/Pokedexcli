package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/danilovict2/Pokedexcli/internal/pokeapi"
)	


func commandMapf(conf *config) error {
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

func commandMapb(conf *config) error {
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

func commandHelp(conf *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	
	commands := getCommands()

	for command := range commands {
		fmt.Println(commands[command].name + ": " + commands[command].description)
	}

	return nil
}

func commandExit(conf *config) error {
	os.Exit(0)
	return nil
}
