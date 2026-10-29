package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range cfg.Command {
		fmt.Printf("%v: %v", cmd.name, cmd.description)
		fmt.Println()
	}
	return nil
}
