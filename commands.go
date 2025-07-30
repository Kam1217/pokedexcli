package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getComands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
