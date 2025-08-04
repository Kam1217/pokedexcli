package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Kam1217/pokedexcli/internal/cache"
)

type Config struct {
	Next     string
	Previous string
	Cache    *cache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func cleanInput(text string) []string {
	lowerStr := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(lowerStr)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "After using 'map', 'explore, to see a list of all the Pokemon in a given location",
			callback:    commandExplore,
		},
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
	cach := cache.NewCache(10 * time.Minute)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	conf := &Config{Cache: cach}
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
			err := cmd.callback(conf, input[1:])
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
