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
	return true, nil
}
