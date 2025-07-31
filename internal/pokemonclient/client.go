package pokemonclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	BaseURL string
	client  *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://pokeapi.co/api/v2",
		client:  &http.Client{},
	}
}

func (c *Client) GetLocationAreas(overrideURL string) (*LocationAreaResponse, error) {
	url := c.BaseURL + "/location-area"
	if overrideURL != "" {
		url = overrideURL
	}
	res, err := c.client.Get(url)
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
