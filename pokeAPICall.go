package main

import (
	"io"
	"net/http"
)

func pokeAPICall(url string) ([]byte, error) {
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
