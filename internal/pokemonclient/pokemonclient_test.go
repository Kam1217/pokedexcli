package pokemonclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	clt := NewClient()
	if clt == nil {
		t.Errorf("No new client created")
	}
	if clt.BaseURL != "https://pokeapi.co/api/v2" {
		t.Errorf("Expected based url to be:https://pokeapi.co/api/v2, but got: %s", clt.BaseURL)
	}
	if clt.client == nil {
		t.Errorf("The http clint has not been initialised")
	}
}

func TestGetLocationAreas(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedJSON := `{
		"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		"previous": null,
		"result":[
		{
			"name": "mock-area-1",
			"URL":"https://pokeapi.co/api/v2/location-area/1/"
		},
		{
			"name": "mock-area-2",
			"URL": "https://pokeapi.co/api/v2/location-area/2/"
		}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedJSON))
	}))

	defer mockServer.Close()

	clt := NewClient()
	clt.BaseURL = mockServer.URL

	res, err := clt.GetLocationAreas("")
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if res.Next != "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" {
		t.Errorf("Expected next URL but got: %s", res.Next)
	}

	if res.Previous != "" {
		t.Errorf("Expected previous URL to be null but got: %s", res.Previous)
	}

	if len(res.Results) != 2 {
		t.Errorf("Expected 2 results but got: %d", len(res.Results))
	}

	if res.Results[0].Name != "mock-area-1" {
		t.Errorf("Expected reult name to be mock-area-1, but got: %s", res.Results[0].Name)
	}

}
