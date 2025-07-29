package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadData(cfg *config) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home file path")
		return
	}
	path := filepath.Join(home, "Documents", "pokedex.json")
	storedData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("There is no pervious save data...")
		return
	}
	data := make(map[string]PokemonInfo)
	if err := json.Unmarshal(storedData, &data); err != nil {
		fmt.Println("Save data corrupted...")
		return
	}
	cfg.Storage = data
}

func WriteData(cfg *config) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home file path")
		return
	}
	path := filepath.Join(home, "Documents", "pokedex.json")
	fmt.Println("Writing to " + path)
	data, err := json.Marshal(cfg.Storage)
	if err != nil {
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
