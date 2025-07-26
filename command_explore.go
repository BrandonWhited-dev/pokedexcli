package main

import (
	"fmt"
	"net/http"
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

}
