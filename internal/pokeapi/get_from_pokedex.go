package pokeapi

import (
	"errors"
)

func (c *Client) GetFromPokedex(name string) (Pokemon, error) {
	if pokemon, exists := c.pokeStorage[name]; exists {
		return pokemon, nil
	} else {
		err := errors.New("you have not caught that pokemon")
		return Pokemon{}, err
	}
}

func (c *Client) GetAllFromPokedex() ([]string, error) {
	numOfItems := len(c.pokeStorage)
	if numOfItems == 0 {
		return []string{}, errors.New("no pokemon in your pokedex")
	}
	pokemonList := make([]string, 0)
	for _, pokemon := range c.pokeStorage {
		pokemonList = append(pokemonList, pokemon.Name)
	}
	return pokemonList, nil
}
