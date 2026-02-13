package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vosram/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient 				pokeapi.Client
	nextLocationsURL			*string
	prevLocationsURL 			*string
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		if command, ok := getCommands()[words[0]]; ok {
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
