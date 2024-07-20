package main

import (
	"fmt"
)

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
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

func main() {
	var input string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		if command, ok := commands[input]; !ok {
			continue
		} else {
			command.callback()

			if command.name == "exit" {
				break
			}
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting Pokedex...")
	return nil
}
