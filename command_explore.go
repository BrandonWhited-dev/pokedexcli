package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type EncountersResponse struct {
	Encounter []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func commandExplore(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[1]
	data := EncountersResponse{}
	if args[1] == "" {
		return errors.New("You need to enter a location. ex: explore canalave-city-area")
	}

	cacheData, exists := cfg.Cache.Get(url)

	if exists {
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
			return errors.New("Area was Not Found. Please try again with a valid area.")
		}
	}

	//Print data
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range data.Encounter {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	//return data
	return nil
}
