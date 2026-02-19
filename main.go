package main

import (
	"time"

	"github.com/vosram/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(conf)
}
