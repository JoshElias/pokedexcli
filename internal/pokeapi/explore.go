package pokeapi

import (
	"encoding/json"
	"fmt"
)

func ExploreRequest(id string) ([]PokemonEncounters, error) {
	url := fmt.Sprintf(
		"%s/%s",
		POKEAPI_LOCATION_URL,
		id,
	)
	data, exists := cache.Get(url)
	var err error
	if !exists {
		data, err = GetRequest(url)
		if err != nil {
			return nil, err
		}
		cache.Add(url, data)
	}

	res := LocationArea{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res.PokemonEncounters, nil
}

func Explore(id string) error {
	encounters, err := ExploreRequest(id)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
