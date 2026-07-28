package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemonDetails(name string) (PokemonDetailsResponse, error) {
	url := baseUrl + "/pokemon/" + name
	var data []byte

	data, isCached := c.cache.Get(url)
	if !isCached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonDetailsResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonDetailsResponse{}, err
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonDetailsResponse{}, err
		}

		if res.StatusCode > 399 {
			return PokemonDetailsResponse{}, fmt.Errorf("bad status code %d\n", res.StatusCode)
		}

		c.cache.Add(url, data)
	}

	response := PokemonDetailsResponse{}

	err := json.Unmarshal(data, &response)
	if err != nil {
		return PokemonDetailsResponse{}, err
	}

	return response, nil
}
