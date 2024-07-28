package pokeapi

import (
	"fmt"
	"math/rand"
)

func Catch(pokemonName string) (bool, error) {
	fmt.Println("Throwing a Pokeball at ", pokemonName, "...")
	pokemon, err := GetPokemon(pokemonName)
	if err != nil {
		return false, err
	}

	randomNumber := rand.Float64()
	threshold := 1.0 / float64(pokemon.BaseExperience+1) * 100
	if randomNumber > threshold {
		fmt.Println(pokemonName, " escape!")
		return false, nil
	}

	PokemonBox[pokemonName] = pokemon
	fmt.Println(pokemonName, " was caught!")
	return true, nil
}

func Inspect(pokemonName string) error {
	pokemon, exists := PokemonBox[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	stats := make(map[string]int)
	for _, stat := range pokemon.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for key, val := range stats {
		fmt.Println("-", key, ": ", val)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println("- ", t.Type.Name)
	}
	return nil
}
