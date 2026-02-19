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