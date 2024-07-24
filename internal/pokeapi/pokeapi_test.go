package pokeapi

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	if err := Map(); err != nil {
		fmt.Println("error printing next maps")
	}
}
