package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		if command, ok := Commands[input]; !ok {
			continue
		} else {
			command.callback()

			if command.name == "exit" {
				break
			}
		}

	}
}
