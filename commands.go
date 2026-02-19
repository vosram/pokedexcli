package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
	return commands
}

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *config, args ...string) error {
	locationRes, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationRes.Next
	conf.prevLocationsURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(conf *config, args ...string) error {
	if conf.prevLocationsURL == nil {
		return fmt.Errorf("You're on the first page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}
	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	name := args[0]
	exploreAreaResp, err := conf.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon: \n", name)
	for _, encounter := range exploreAreaResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("a name of a pokemon is needed")
	}
	pokemonName := args[0]
	pokemon, err := conf.pokeapiClient.FetchPokemon(pokemonName)	

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	// Ceiling for probablity should be 700
	didCatch := getProbability(700, pokemon.BaseExperience)
	if didCatch == false {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	conf.pokeapiClient.StorePokemon(pokemon)
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}

func getProbability(ceil, target int) bool {
	randomNum := rand.Intn(ceil)
	fmt.Printf("random num is: %d and target is:  %d\n", randomNum, target)
	if randomNum >= target {
		return true
	}
	return false
}