package pokeapi

import (
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
