package pokeapi

import (
	"net/http"
	"time"

	"github.com/ingodinho/pokedex_cli/internal/pokecache"
)

const baseUrl string = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      *pokecache.PokeCache
}

func NewClient(cacheTtl time.Duration) *Client {
	cache := pokecache.NewPokeCache(cacheTtl)

	return &Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: cache,
	}
}
