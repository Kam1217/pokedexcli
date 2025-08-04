package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/Kam1217/pokedexcli/internal/cache"
	"github.com/Kam1217/pokedexcli/internal/pokemonclient"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "  hello world   ",
		expected: []string{"hello", "world"},
	},
		{
			input:    "I am making a pokedex   ",
			expected: []string{"i", "am", "making", "a", "pokedex"},
		},
		{
			input:    "I like Mew and Jigglypuff",
			expected: []string{"i", "like", "mew", "and", "jigglypuff"},
		}}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual slice: %d, does not match length of expected slice: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word %s, does not match expected word %s", word, expectedWord)
			}
		}

	}
}

func TestCommandHelp(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := commandHelp(nil, nil)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)
	outputStr := strings.TrimSpace(string(output))
	expectedStrings := []string{
		"catch: Catch a specific pokemon by calling 'catch' with the pokemon name",
		"explore: After using 'map', 'explore, to see a list of all the Pokemon in a given location",
		"map: Displays the names of 20 location areas in the Pokemon world",
		"mapb: Displays the names of 20 previous location areas in the Pokemon world",
		"help: Displays a help message",
		"exit: Exit the Pokedex",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(outputStr, expected) {
			t.Errorf("Expected output to contain %s", expected)
		}
	}

	if err != nil {
		t.Errorf("Expected no error but got: \n%v", err)
	}
}

func TestCommandExit(t *testing.T) {
	if os.Getenv("BE_EXIT") == "0" {
		commandExit(nil, nil)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCommandExit")
	cmd.Env = append(os.Environ(), "BE_EXIT=0")
	err := cmd.Run()
	if err == nil {
		return
	}

	t.Fatalf("process run with err %v, want exit status 0", err)
}

func TestCommandMap(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedJSON := `{
		"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		"previous": null,
		"results":[
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

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cach := cache.NewCache(10 * time.Minute)
	conf := &Config{
		Next:          mockServer.URL,
		Previous:      "",
		Cache:         cach,
		PokemonClient: pokemonclient.NewClient(cach),
	}

	err := commandMap(conf, nil)
	w.Close()
	output, _ := io.ReadAll(r)
	os.Stdout = oldStdout
	outputStr := strings.TrimSpace(string(output))

	if err != nil {
		t.Error("Expected no error but go: %w", err)
	}

	expectedNames := []string{"mock-area-1", "mock-area-2"}
	for _, name := range expectedNames {
		if !strings.Contains(outputStr, name) {
			t.Errorf("Expected output to contain %s", name)
		}
	}

	if conf.Next != "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" {
		t.Errorf("Expected Next to be updated")
	}
}

func TestCommandMapb(t *testing.T) {
	t.Run("Normal operation", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedJSON := `{
		"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		"previous":"https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		"results":[
		{
			"name": "mock-area-3",
			"URL":"https://pokeapi.co/api/v2/location-area/3/"
		}
			]
		}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedJSON))
		}))
		defer mockServer.Close()

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		cach := cache.NewCache(10 * time.Minute)
		conf := &Config{
			Next:          "",
			Previous:      mockServer.URL,
			Cache:         cach,
			PokemonClient: pokemonclient.NewClient(cach),
		}

		err := commandMapb(conf, nil)
		w.Close()
		output, _ := io.ReadAll(r)
		os.Stdout = oldStdout
		outputStr := strings.TrimSpace(string(output))

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if !strings.Contains(outputStr, "mock-area-3") {
			t.Errorf("Expected output to contain: mock-area-3")
		}
	})

	t.Run("First page error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedJSON := `{
		"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		"previous": null,
		"results": []
		}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedJSON))
		}))

		defer mockServer.Close()

		cach := cache.NewCache(10 * time.Minute)
		conf := &Config{
			Next:          "",
			Previous:      mockServer.URL,
			Cache:         cach,
			PokemonClient: pokemonclient.NewClient(cach),
		}

		err := commandMapb(conf, nil)

		if err == nil {
			t.Errorf("Expected first page error but got nil")
		}

		expectedError := "you're on the first page"
		if err.Error() != expectedError {
			t.Errorf("Expected %s, but got %s", expectedError, err.Error())
		}
	})
}

func TestCommandExplore(t *testing.T) {
	t.Run("Normal operation", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedJSON := `
        {
            "pokemon_encounters": [
                {
                    "pokemon": {
                        "name": "mock-name-1",
                        "url": "https://pokeapi.co/api/v2/pokemon/1/"
                    }
                },
                {
                    "pokemon": {
                        "name": "mock-name-2",
                        "url": "https://pokeapi.co/api/v2/pokemon/2/"
                    }
                }
            ]
        }`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedJSON))
		}))
		defer mockServer.Close()

		cach := cache.NewCache(10 * time.Minute)
		newClient := pokemonclient.NewClient(cach)
		newClient.BaseURL = mockServer.URL

		conf := &Config{
			Next:          mockServer.URL,
			Previous:      mockServer.URL,
			Cache:         cach,
			PokemonClient: newClient,
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := commandExplore(conf, []string{"mock-area"})

		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !strings.Contains(string(out), "mock-name-1") {
			t.Errorf("output did not contain mock-name-1")
		}
		if !strings.Contains(string(out), "mock-name-2") {
			t.Errorf("output did not contain mock-name-2")
		}
	})

	t.Run("Empty argument list", func(t *testing.T) {
		cach := cache.NewCache(10 * time.Minute)
		client := pokemonclient.NewClient(cach)

		conf := &Config{
			PokemonClient: client,
			Cache:         cach,
		}

		err := commandExplore(conf, []string{})

		if err == nil {
			t.Errorf("expected error for empty args, got nil")
		}
		wantMsg := "pokemon location area cannot be empty"
		if err != nil && err.Error() != wantMsg {
			t.Errorf("expected error '%s', got '%v'", wantMsg, err)
		}
	})
}
