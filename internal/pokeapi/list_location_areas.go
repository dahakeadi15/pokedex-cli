package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dahakeadi15/pokedex-cli/internal/pokecache"
)

func (c *Client) ListLocationAreas(pageURL *string, cache pokecache.Cache) (ApiRespLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ApiRespLocationAreas{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ApiRespLocationAreas{}, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return ApiRespLocationAreas{}, err
		}
		data = body
		cache.Add(url, data)
	}

	locationAreasResponse := ApiRespLocationAreas{}
	if err := json.Unmarshal(data, &locationAreasResponse); err != nil {
		return ApiRespLocationAreas{}, err
	}

	return locationAreasResponse, nil
}
