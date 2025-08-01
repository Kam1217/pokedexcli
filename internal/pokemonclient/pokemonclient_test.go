package pokemonclient

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
	"github.com/Kam1217/pokedexcli/internal/cache"
	"time"
)

func TestNewClient(t *testing.T) {
	cash := cache.NewCache(10 * time.Minute)
	clt := NewClient(cash)
	if clt == nil {
		t.Fatalf("No new client created")
	}
	if clt.BaseURL != "https://pokeapi.co/api/v2" {
		t.Fatalf("Expected based url to be:https://pokeapi.co/api/v2, but got: %s", clt.BaseURL)
	}
	if clt.client == nil {
		t.Fatalf("The http clint has not been initialised")
	}
}

func TestGetLocationAreas(t *testing.T) {
	//TODO: USE THIS FOR COMMAND TESTS
	// mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	expectedJSON := `{
	// 	"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
	// 	"previous": null,
	// 	"results":[
	// 	{
	// 		"name": "mock-area-1",
	// 		"URL":"https://pokeapi.co/api/v2/location-area/1/"
	// 	},
	// 	{
	// 		"name": "mock-area-2",
	// 		"URL": "https://pokeapi.co/api/v2/location-area/2/"
	// 	}
	// 		]
	// 	}`
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(expectedJSON))
	// }))

	// defer mockServer.Close()
	cach := cache.NewCache(10 * time.Minute)
	clt := NewClient(cach)
	clt.BaseURL = "https://pokeapi.co/api/v2"
	res, err := clt.GetLocationAreas("")
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if res.Next != "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" {
		t.Fatalf("Expected next URL but got: %s", res.Next)
	}

	if res.Previous != "" {
		t.Fatalf("Expected previous URL to be null but got: %s", res.Previous)
	}

	if len(res.Results) != 20 {
		t.Logf("%v", res.Results)
		t.Fatalf("Expected 20 results but got: %d", len(res.Results))
	}

}
