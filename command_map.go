package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LocationAreaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandMap(cfg *config, args []string) error {
	url := cfg.Next
	data := LocationAreaResponse{}
	if url == "" {
		return errors.New("There is no next data.")
	}

	cacheData, exists := cfg.Cache.Get(url)

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
			return err
		}
	}

	//set config
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	//Print data
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	//return data
	return nil
}

func commandMapb(cfg *config, args []string) error {
	url := cfg.Previous
	data := LocationAreaResponse{}
	if url == "" {
		return errors.New("There is no previous data.")
	}
	cacheData, exists := cfg.Cache.Get(url)

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
			return err
		}
	}

	//set config
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	//Print data
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	//return data
	return nil
}
