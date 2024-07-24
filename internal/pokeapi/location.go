package pokeapi

type LocationArea struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	GameIndex            int                    `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location             Location               `json:"location"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
