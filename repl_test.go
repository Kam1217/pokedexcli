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
			expected: []string{"i", "am", "making", "a", "pokedex"},
		},
		{
			input:    "I like Mew and Jigglypuff",
			expected: []string{"i", "like", "mew", "and", "jigglypuff"},
		}}

		for _,c := range cases{
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected){
				t.Errorf("length of actual slice: %d, does not match length of expected slice: %d", len(actual), len(c.expected))
			}

			for i := range actual{
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord{
					t.Errorf("Word %s, does not match expected word %s", word, expectedWord)
				}
			}

		}
}
