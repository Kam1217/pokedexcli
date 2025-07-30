package main

type LocationAreaResponse struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results [] struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`  
}

