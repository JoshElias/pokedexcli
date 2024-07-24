package pokeapi

const POKEAPI_URL = "https://pokeapi.co/api/v2/location-area"

type Version struct {
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

type PokeResponseList struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next"`
	Previous string                 `json:"previous"`
	Results  []PokeResponseListItem `json:"results"`
}

type PokeResponseListItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
