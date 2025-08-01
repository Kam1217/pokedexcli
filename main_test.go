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

	err := commandHelp(nil)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)
	outputStr := strings.TrimSpace(string(output))
	expectedStrings := []string{
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
		commandExit(nil)
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
		Next:     mockServer.URL,
		Previous: "",
		Cache:    cach,
	}

	err := commandMap(conf)
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
			Next:     "",
			Previous: mockServer.URL,
			Cache:    cach,
		}

		err := commandMapb(conf)
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
			Next:     "",
			Previous: mockServer.URL,
			Cache:    cach,
		}

		err := commandMapb(conf)

		if err == nil {
			t.Errorf("Expected first page error but got nil")
		}

		expectedError := "you're on the first page"
		if err.Error() != expectedError {
			t.Errorf("Expected %s, but got %s", expectedError, err.Error())
		}
	})
}
