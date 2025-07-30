package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error		
}

func cleanInput(text string) []string {

	lowerStr := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(lowerStr)

	return words
}

func commandExit() error {
	fmt.Print("Closing the Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex!")
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		args := cleanInput(input)

		if len(input) == 0{
			fmt.Print("Pokedex > ")
			continue
		}
		
		_, ok := commands[args[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		
		fmt.Print("Pokedex > ")
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
