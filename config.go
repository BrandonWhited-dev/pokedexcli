package main

import "github.com/brandonwhited-dev/pokedexcli/internal/pokecache"

type config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Storage  map[string]PokemonInfo
}
