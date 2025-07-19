package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := cfg.Next
	if url == "" {
		return errors.New("There is no next data.")
	}
	// make the get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//get the response
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	//read data
	data := LocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
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

func commandMapb(cfg *config) error {
	url := cfg.Previous
	if url == "" {
		return errors.New("There is no previous data.")
	}
	// make the get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//get the response
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	//read data
	data := LocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
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
