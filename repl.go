package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	
	"github.com/danilovict2/Pokedexcli/internal/pokecache"
)

type config struct {
	Previous *string
	Next *string
	Cache pokecache.Cache
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	conf := &config{
		Cache: pokecache.NewCache(5 * time.Second),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		comm, exists := commands[words[0]]
		
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		
		err := comm.callback(conf)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}