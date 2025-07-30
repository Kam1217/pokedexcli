package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	lowerStr := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(lowerStr)
	return words
}

func getComands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmdArg := input[0]

		cmd, ok := getComands()[cmdArg]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
	}

}
