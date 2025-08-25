package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ApiRespLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, exists := c.cache.Get(url)
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

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return ApiRespLocationAreas{}, err
		}

		c.cache.Add(url, data)
	}

	locationAreasResponse := ApiRespLocationAreas{}
	if err := json.Unmarshal(data, &locationAreasResponse); err != nil {
		return ApiRespLocationAreas{}, err
	}

	return locationAreasResponse, nil
}
