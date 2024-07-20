package main

import (
	"fmt"
)

var Commands = map[string]cliCommand{}

func init() {
	RegisterCommand("help", "Displays a help message", commandHelp)
	RegisterCommand("exit", "Exit the Pokedex", commandExit)
}

func RegisterCommand(name string, description string, commandFunc func() error) {
	if _, ok := Commands[name]; ok {
		fmt.Println("overriding existing command: ", name)
	}
	Commands[name] = cliCommand{
		name:        name,
		description: description,
		callback:    commandFunc,
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
