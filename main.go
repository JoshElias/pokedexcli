package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
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
