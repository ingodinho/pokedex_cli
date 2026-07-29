package pokeapi

import "sync"

type Pokedex struct {
	mu      sync.Mutex
	pokemon map[string]PokemonDetailsResponse
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemon: make(map[string]PokemonDetailsResponse),
	}
}

func (pd *Pokedex) Add(pm PokemonDetailsResponse) {
	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.pokemon[pm.Name] = pm
}

func (pd *Pokedex) Get(key string) (PokemonDetailsResponse, bool) {
	pd.mu.Lock()
	defer pd.mu.Unlock()

	pokemon, isCached := pd.pokemon[key]

	return pokemon, isCached
}
