package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
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

	expectedOutput := `Usage:

help: Displays a help message
exit: Exit the Pokedex`

	if outputStr != expectedOutput {
		t.Errorf("Expected output: \n%s, \nActual output: \n%s", expectedOutput, outputStr)
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
