package pokeapi

import (
	"encoding/json"
	"io"
)

const baseUrl string = "https://pokeapi.co/api/v2"

func FetchLocationAreas(next string) (LocationAreaResponse, error){
	c := NewClient()

	url := baseUrl + "/location-area"
	if next != "" {
		url = next
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}
