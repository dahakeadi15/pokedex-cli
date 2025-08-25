package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInArea(areaName string) ([]string, error) {
	url := baseURL + "/location-area/" + areaName

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if string(data) == "Not Found" {
			return nil, fmt.Errorf("pokemon not found in area: %s", areaName)
		}

		c.cache.Add(url, data)
	}

	locAreaByIDRes := ApiRespLocationAreaByName{}
	if err := json.Unmarshal(data, &locAreaByIDRes); err != nil {
		fmt.Println(string(data))
		return nil, fmt.Errorf("error unmarshaling: %w", err)
	}

	pokemonList := []string{}
	for _, encounter := range locAreaByIDRes.PokemonEncounters {
		pokemonList = append(pokemonList, encounter.Pokemon.Name)
	}

	return pokemonList, nil
}
