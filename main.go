package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		words := strings.Fields(input)
		command, ok := Commands[words[0]]
		if !ok {
			continue
		}
		if command.name == "exit" {
			break
		}

		if len(words) > 1 {
			command.callback(words[1:]...)
		} else {
			command.callback()
		}
	}
}
