package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaListResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	return GETRequestWithCache[LocationAreaListResponse](c, url)
}

func (c *Client) GetLocationArea(locationName string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + locationName

	return GETRequestWithCache[LocationAreaResponse](c, url)
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	return GETRequestWithCache[Pokemon](c, url)
}

func GETRequestWithCache[T any](c *Client, url string) (T, error) {
	var zero T

	dat, found := c.cache.Get(url)
	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return zero, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return zero, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return zero, err
		}

		c.cache.Add(url, body)
		dat = body
	}

	var resp T
	err := json.Unmarshal(dat, &resp)
	if err != nil {
		return zero, err
	}

	return resp, nil
}
