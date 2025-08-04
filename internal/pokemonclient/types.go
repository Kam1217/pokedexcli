package pokemonclient

type LocationAreaResponse struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`  
}

type FindPokemonResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}


type CatchPokemonResponse struct{
	BaseExperience int `json:"base_experience"`
	Forms [] struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"forms"`
}
