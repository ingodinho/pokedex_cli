package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocationAreas(givenUrl *string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area/?offset=0&limit=20"
	if givenUrl != nil {
		url = *givenUrl
	}

	var data []byte

	data, isCached := c.cache.Get(url)

	if !isCached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		defer res.Body.Close()

		if res.StatusCode > 399 {
			return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		c.cache.Add(url, data)
	}

	locationAreaResponse := LocationAreaResponse{}
	err := json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}
