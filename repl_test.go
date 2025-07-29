package main

import (
	"testing"
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
			expected: []string{"I", "am", "making", "a", "pokedex"},
		},
		{
			input:    "I like Mew and Jigglypuff",
			expected: []string{"I", "like", "Mew", "and", "Jigglypuff"},
		}}
}
