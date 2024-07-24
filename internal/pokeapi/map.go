package pokeapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

const firstMapUrl = POKEAPI_URL + "?offset=0&limit=20"

var nextMapUrl = strings.Clone(firstMapUrl)
var prevMapUrl string

func MapRequest(url string) (PokeResponseList, error) {
	data, err := GetRequest(url)
	if err != nil {
		return PokeResponseList{}, err
	}

	res := PokeResponseList{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return PokeResponseList{}, err
	}
	return res, nil
}

func HandleMap(url string) error {
	list, err := MapRequest(url)
	if err != nil {
		return err
	}
	nextMapUrl = list.Next
	prevMapUrl = list.Previous
	for _, location := range list.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func Map() error {
	if nextMapUrl == "" {
		return fmt.Errorf("you've reached the end of the location area list")
	}
	return HandleMap(nextMapUrl)
}

func MapB() error {
	if prevMapUrl == "" {
		return fmt.Errorf("you're at the beginning of the location area list")
	}
	return HandleMap(prevMapUrl)
}
