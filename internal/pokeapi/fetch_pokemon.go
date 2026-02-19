package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	
	// cache hit
	if cacheData, cacheExists := c.cache.Get(url); cacheExists {
		var pokemon Pokemon
		err := json.Unmarshal(cacheData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	
	// cache miss
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return  Pokemon{}, errors.New("pokemon fetch not 200 OK")
	}
	dat, err := io.ReadAll(res.Body)	
	if err != nil {
		return Pokemon{}, err
	}
	
	var pokemon Pokemon
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemon, nil
}

func (c *Client) StorePokemon(pokemon Pokemon) {
	c.pokeStorage[pokemon.Name] = pokemon
}