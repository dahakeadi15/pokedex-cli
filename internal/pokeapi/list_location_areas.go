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

	locationAreasResponse := ApiRespLocationAreas{}
	if err := json.Unmarshal(body, &locationAreasResponse); err != nil {
		return ApiRespLocationAreas{}, err
	}

	return locationAreasResponse, nil
}
