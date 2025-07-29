package main

import (
	"github.com/brandonwhited-dev/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	configuration := &config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Cache:    pokecache.NewCache(10 * time.Second),
		Storage:  make(map[string]PokemonInfo),
	}
	replStart(configuration)
}
