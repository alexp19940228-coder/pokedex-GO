package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]
		if command, exists := commands[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
		} else {
			fmt.Printf("Unknown command: '%s'. Type 'help' for a list of commands.\n", commandName)
		}

	}

}
