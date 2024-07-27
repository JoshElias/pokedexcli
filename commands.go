package main

import (
	"errors"
	"fmt"
	"internal/pokeapi"
)

var Commands = map[string]cliCommand{}

func init() {
	RegisterCommand("help", "Displays a help message", commandHelp)
	RegisterCommand("exit", "Exit the Pokedex", commandExit)
	RegisterCommand("map", "List the next 20 locations", commandMap)
	RegisterCommand("mapb", "List the previous 20 locations", commandMapB)
	RegisterCommand("explore", "List Pokemon encounters for a location", commandExplore)
}

func RegisterCommand(name string, description string, commandFunc func(...string) error) {
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
	callback    func(...string) error
}

func commandHelp(args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandExit(args ...string) error {
	fmt.Println("Exiting Pokedex...")
	return nil
}

func commandMap(args ...string) error {
	return pokeapi.Map()
}

func commandMapB(args ...string) error {
	return pokeapi.MapB()
}

func commandExplore(args ...string) error {
	if len(args) == 0 {
		return errors.New("explore command needs a location area name or id")
	}
	return pokeapi.Explore(args[0])
}
