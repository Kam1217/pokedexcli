package main

import (
	"fmt"
	"os"

	"github.com/Kam1217/pokedexcli/internal/pokemonclient"
)

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *Config) error {
	c := pokemonclient.NewClient()
	res, err := c.GetLocationAreas(conf.Next)
	if err != nil {
		return err
	}
	conf.Next = res.Next
	conf.Previous = res.Previous
	for _, location := range res.Results {
		println(location.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	c := pokemonclient.NewClient()
	res, err := c.GetLocationAreas(conf.Previous)
	if err != nil {
		return err
	}
	conf.Previous = res.Previous
	conf.Next = res.Next
	for _, location := range res.Results {
		println(location.Name)
	}
	return nil
}
