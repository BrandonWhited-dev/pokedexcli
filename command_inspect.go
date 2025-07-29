package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 2 || args[1] == "" {
		return errors.New("Please enter a pokemons name. ex: catch Pikachu")
	}
	pokemonName := args[1]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	cacheData, exists := cfg.Cache.Get(url)

	data := PokemonInfo{}

	if exists {
		//decode cache data
		if err := json.Unmarshal(cacheData, &data); err != nil {
			return err
		}
	} else {
		body, err := pokeAPICall(url)
		if err != nil {
			return err
		}
		//cache data
		cfg.Cache.Add(url, body)

		//decode body
		if err := json.Unmarshal(body, &data); err != nil {
			return errors.New("Pokemon not found. Please enter a valid name.")
		}
	}

	PrintInspect(data)

	return nil
}

func PrintInspect(info PokemonInfo) {
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Base EXP: %d\n", info.Exp)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Println("Stats:")
	for _, pokeStats := range info.Stats {
		fmt.Printf("   -%s: %d\n", pokeStats.Stat.Name, pokeStats.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range info.Types {
		fmt.Printf("   -%s\n", pokeType.Types.Name)
	}
}
