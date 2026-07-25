package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseUrl string = "https://pokeapi.co/api/v2"

func FetchLocationAreas(next string) (LocationAreaResponse, error){
	url := baseUrl + "/location-area"
	if next != "" {
		url = next
	}

	res, err := http.Get(url)
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
