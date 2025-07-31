package pokemonClient

type LocationAreaResponse struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results [] struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`  
}

