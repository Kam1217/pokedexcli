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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Pokedex >")
		arg := scanner.Text()
		argWords := cleanInput(arg)
		fmt.Println("Your command was:", argWords[0])
	}
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
    }
}
