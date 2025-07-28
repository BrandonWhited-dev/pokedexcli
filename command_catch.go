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

