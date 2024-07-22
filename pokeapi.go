package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const POKEAPI_URL = "https://pokeapi.co/api/v2/location-area"

type PokeAPIList struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []PokeAPIListItem `json:"results"`
}

type PokeAPIListItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	GameIndex            int                    `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location             Location               `json:"location"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodVersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod                 `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetails `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Names struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetails struct {
	MinLevel        int    `json:"min_level"`
	MaxLevel        int    `json:"max_level"`
	ConditionValues []any  `json:"condition_values"`
	Chance          int    `json:"chance"`
	Method          Method `json:"method"`
}

type EncounterVersionDetails struct {
	Version          Version            `json:"version"`
	MaxChance        int                `json:"max_chance"`
	EncounterDetails []EncounterDetails `json:"encounter_details"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon                   `json:"pokemon"`
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}

const firstMapUrl = POKEAPI_URL + "?offset=0&limit=20"

var nextMapUrl = strings.Clone(firstMapUrl)
var prevMapUrl string

func MapRequest(url string) (PokeAPIList, error) {
	data, err := GetRequest(url)
	if err != nil {
		return PokeAPIList{}, err
	}

	res := PokeAPIList{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return PokeAPIList{}, err
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
