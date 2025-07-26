package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func pokemapCall(url string) ([]byte, error) {

	// make the get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//get the response
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//read data
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
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
		body, err := pokemapCall(url)
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
		body, err := pokemapCall(url)
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
