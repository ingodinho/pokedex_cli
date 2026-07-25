package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationAreaResponse struct {
	PokeResponse
	Results []LocationArea `json:"results"`
}
