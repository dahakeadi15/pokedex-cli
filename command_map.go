package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiResponse struct {
	Count    uint
	Next     string
	Previous any
	Results  []Location
}

type Location struct {
	Name string
	URL  string
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location")
	if err != nil {
		return fmt.Errorf("error fetching locations: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("error fetching locations: api failed with code %d and body - %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	apiRes := ApiResponse{}
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return fmt.Errorf("error parsing data: %w", err)
	}

	for _, location := range apiRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
