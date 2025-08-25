package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (ApiRespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ApiRespPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ApiRespPokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return ApiRespPokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokemonResp := ApiRespPokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return ApiRespPokemon{}, nil
	}
	return pokemonResp, nil
}
