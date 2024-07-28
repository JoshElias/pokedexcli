package pokeapi

import (
	"encoding/json"
	"fmt"
)

const POKEAPI_POKEMON_URL = POKEAPI_URL + "/pokemon"

var PokemonBox = make(map[string]Pokemon)

func Catch(name string) error {
	return nil
}

func GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf(
		"%s/%s",
		POKEAPI_POKEMON_URL,
		name,
	)
	data, exists := cache.Get(url)
	var err error
	if !exists {
		data, err = GetRequest(url)
		if err != nil {
			return Pokemon{}, err
		}
		cache.Add(url, data)
	}

	res := Pokemon{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return Pokemon{}, err
	}
	return res, nil
}
