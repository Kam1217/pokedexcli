package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func cleanInput(text string) []string {
	lowerStr := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(lowerStr)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world",
			callback:    commandMapb,
		},
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
	conf := &Config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmdArg := input[0]

		cmd, ok := getCommands()[cmdArg]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := cmd.callback(conf)
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
