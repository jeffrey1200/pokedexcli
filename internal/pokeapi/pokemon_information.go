package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonInformation, error) {
	url := baseApiUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonData := PokemonInformation{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return PokemonInformation{}, nil
		}

		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInformation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInformation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInformation{}, err
	}

	pokemonData := PokemonInformation{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return PokemonInformation{}, err
	}

	c.cache.Add(url, data)

	return pokemonData, nil
}
