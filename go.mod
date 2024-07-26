module github.com/JoshElias/pokedexcli

go 1.22.5

replace internal/pokeapi v0.0.1 => ./internal/pokeapi

replace internal/pokecache v0.0.1 => ./internal/pokecache

require internal/pokeapi v0.0.1

require internal/pokecache v0.0.1 // indirect
