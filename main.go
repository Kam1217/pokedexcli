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
	fmt.Println("Closing the Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println("")
	return nil
}

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex!")
	for scanner.Scan() {
		input := scanner.Text()
		args := cleanInput(input)

		if len(input) == 0{
			continue
		}
		
		cmd, ok := commands[args[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback() 
		if err != nil{
			fmt.Println("Error: ", err.Error())
			continue
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
