package internal

type LocationAreaResponse struct {
	PokeResponse
	Results []LocationArea `json:"results"`
}

type PokeResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
