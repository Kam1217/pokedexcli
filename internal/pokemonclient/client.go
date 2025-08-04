package pokemonclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Kam1217/pokedexcli/internal/cache"
)

type Client struct {
	BaseURL string
	client  *http.Client
	Cache   *cache.Cache
}

func NewClient(cash *cache.Cache) *Client {
	return &Client{
		BaseURL: "https://pokeapi.co/api/v2",
		client:  &http.Client{},
		Cache:   cash,
	}
}

func (c *Client) GetLocationAreas(overrideURL string) (*LocationAreaResponse, error) {
	url := c.BaseURL + "/location-area"
	if overrideURL != "" {
		url = overrideURL
	}
	var body []byte

	val, ok := c.Cache.Get(url)
	if ok {
		body = val
	} else {
		res, err := c.client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response: %w", err)
		}

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		c.Cache.Add(url, body)
	}

	var locationResp LocationAreaResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}
	return &locationResp, nil
}

func (c *Client) FindPokemon (overrideURL, name string) (*FindPokemonResponse, error){
	url := c.BaseURL + "/location-area/" + name
	if overrideURL != ""{
		url = overrideURL
	}
	var body []byte
	val, ok := c.Cache.Get(url)
	if ok {
		body = val
	} else {
		res, err := c.client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response: %w", err)
		}

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		c.Cache.Add(url, body)
	}	
	var findPokemonResp FindPokemonResponse	
	
	if err := json.Unmarshal(body, &findPokemonResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}
	return &findPokemonResp, nil	
}