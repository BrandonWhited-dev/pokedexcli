package main

import (
	"fmt"
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

	//read the data and set in in configuration
	fmt.Println("Searching for save data...")
	ReadData(configuration)

	replStart(configuration)
}
