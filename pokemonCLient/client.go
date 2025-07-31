package pokemonClient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func GetLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	var locationResp LocationAreaResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &locationResp, nil
}
