package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, found := c.cache.Get(url)
	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		c.cache.Add(url, body)
		dat = body
	}

	locationAreaResp := LocationAreaResponse{}
	err := json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResp, nil
}
