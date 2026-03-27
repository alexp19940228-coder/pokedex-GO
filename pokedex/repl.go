package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexp19940228-coder/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous string
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List all available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display a map of all Pokemon locations",
			callback:    commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Display a map of all previous Pokemon locations",
			callback:    commandMapBack,
		},
	}

}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	textSplit := strings.Fields(textLower)
	return textSplit
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display a map of all Pokemon locations")
	return nil
}

func commandMap() error {
	// Fetch the Pokemon map data
	pokemonMap := internal.GetPokemonMap()
	if pokemonMap == nil {
		return fmt.Errorf("Failed to fetch Pokemon map data")
	}
	for _, location := range pokemonMap.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack() error {
	pokemonMap := internal.GetPokemonMap()
	if pokemonMap == nil {
		return fmt.Errorf("Failed to fetch Pokemon map data")
	}
	for i := len(pokemonMap.Results) - 1; i >= 0; i-- {
		fmt.Println(pokemonMap.Results[i].Name)
	}
	return nil
}
