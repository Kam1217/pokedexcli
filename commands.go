package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(conf *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, args []string) error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *Config, args []string) error {
	c := conf.PokemonClient
	res, err := c.GetLocationAreas(conf.Next)
	if err != nil {
		return err
	}
	conf.Next = res.Next
	conf.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(conf *Config, args []string) error {
	c := conf.PokemonClient
	res, err := c.GetLocationAreas(conf.Previous)
	if err != nil {
		return err
	}
	conf.Previous = res.Previous
	conf.Next = res.Next
	if conf.Previous == "" {
		return errors.New("you're on the first page")
	}
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(conf *Config, args []string) error {
	c := conf.PokemonClient
	if len(args) == 0 {
		return errors.New("pokemon location area cannot be empty")
	}
	res, err := c.FindPokemon(args[0])
	if err != nil {
		return err
	}
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *Config, args []string) error {
	c := conf.PokemonClient

	if len(args) == 0 {
		return errors.New("pokemon name cannot be empty")
	}

	fmt.Printf("Throwing a Pokeball at %s...", args[0])

	res, err := c.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	randomNum := rand.Intn(res.BaseExperience) + 1
	catchThreshold := 100

	if randomNum <= catchThreshold {
		fmt.Printf("%s was caught!\n", args[0])
		conf.Pokedex[args[0]] = Pokemon{
			Name:           args[0],
			BaseExperience: res.BaseExperience,
		}
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}
	return nil
}
