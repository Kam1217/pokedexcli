package pokemonclient

import(
	"testing"
	// "net/http"
	// "net/http/httptest"
)

func TestNewClient(t *testing.T){
	clt := NewClient()
	if clt == nil {
		t.Errorf("No new client created")
	}

	if clt.BaseURL != "https://pokeapi.co/api/v2"{
		t.Errorf("Expected based url to be:https://pokeapi.co/api/v2, but got: %s", clt.BaseURL)
	}
	if clt.client == nil {
		t.Errorf("The http clint has not been initialised")
	}
}

