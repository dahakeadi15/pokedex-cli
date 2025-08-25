package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokemonResp := Pokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}
	return pokemonResp, nil
}
