package pokeapi

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	if err := Map(); err != nil {
		t.Fatal("error printing next maps")
	}
}

func TestExplore(t *testing.T) {
	const location = "canalave-city-area"
	if err := Explore(location); err != nil {
		t.Fatal("error exploring location")
	}
}

func TestGetPokemon(t *testing.T) {
	const pokemon = "pikachu"
	if pokemon, err := GetPokemon(pokemon); err != nil {
		t.Fatal("error exploring location")
	} else {
		fmt.Println(pokemon.Name)
	}
}
