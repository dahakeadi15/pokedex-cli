package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(areaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + areaName

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}

		c.cache.Add(url, data)
	}

	locationResp := LocationArea{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return LocationArea{}, err
	}

	return locationResp, nil
}
