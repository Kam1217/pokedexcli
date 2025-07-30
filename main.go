package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		arg := scanner.Text()
		argWords := cleanInput(arg)
		fmt.Println("Your command was:", argWords[0])
		fmt.Print("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
