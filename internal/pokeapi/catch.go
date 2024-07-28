package pokeapi

import (
	"fmt"
	"math/rand"
)

func Catch(pokemonName string) (bool, error) {
	pokemon, err := GetPokemon(pokemonName)
	if err != nil {
		return false, err
	}

	fmt.Println("base xp ", pokemon.BaseExperience)
	randomNumber := rand.Float64()
	fmt.Println("random number ", randomNumber)
	threshold := 1.0 / float64(pokemon.BaseExperience+1) * 100
	fmt.Println("threshold ", threshold)
	if randomNumber > threshold {
		fmt.Println(pokemonName, " escape!")
		return false, nil
	}

	PokemonBox[pokemonName] = pokemon
	fmt.Println(pokemonName, " caught!")
	return true, nil
}
