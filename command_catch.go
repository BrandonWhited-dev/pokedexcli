package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PokemonInfo struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Exp       int         `json:"base_experience"`
	Weight    int         `json:"weight"`
	Abilities []Abilities `json:"abilities"`
	Types     []Types     `json:"types"`
	Moves     []Moves     `json:"moves"`
}

type Moves struct {
	Move Move `json:"move"`
}

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Abilities struct {
	Ability Ability `json:"ability"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Types struct {
	Types Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandCatch(cfg *config, args []string) error {
	if len(args) < 2 || args[1] == "" {
		return errors.New("Please enter a pokemons name. ex: catch Pikachu")
	}
	pokemonName := args[1]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	if _, exists := cfg.Storage[pokemonName]; exists {
		return errors.New("You have already caught " + pokemonName)
	}
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

	if catch(data.Name, data.Exp) {
		cfg.Storage[data.Name] = data
	}

	return nil
}

// logic if pokemon is caught
func catch(name string, baseExp int) bool {
	pokeballChance := 75
	//Throwing pokeball message
	fmt.Printf("Throwing a pokeball at %s", name)
	for range 3 {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	fmt.Println("")

	if pokeballChance > rand.Intn(baseExp) {
		fmt.Printf("You caught a %s!\n", name)
		return true
	} else {
		fmt.Printf("%s escaped.\n", name)
		return false
	}
}
