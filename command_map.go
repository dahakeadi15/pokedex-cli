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
	Previous string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	URL  string
}

func commandMap(config *Config) error {
	res, err := http.Get(config.Next)
	if err != nil {
		return fmt.Errorf("error fetching location areas: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("error fetching location areas: api failed with code %d and body - %s", res.StatusCode, body)
	}
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	apiRes := ApiResponse{}
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return fmt.Errorf("error parsing data: %w", err)
	}

	for _, locationArea := range apiRes.Results {
		fmt.Println(locationArea.Name)
	}

	config.Next = apiRes.Next
	config.Previous = apiRes.Previous

	return nil
}
