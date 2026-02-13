package main

import (
	"fmt"
	"os"
)


type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
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
	}
	return commands
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *config) error {
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

func commandMapb(conf *config) error {
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
