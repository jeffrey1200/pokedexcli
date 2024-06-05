package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonsInArea(areaName string) (PokemonsInLocation, error) {

	url := baseApiUrl + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		pokemonInArea := PokemonsInLocation{}
		err := json.Unmarshal(val, &pokemonInArea)
		if err != nil {
			return PokemonsInLocation{}, err
		}

		return pokemonInArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonsInLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonsInLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonsInLocation{}, err
	}

	pokemonsInArea := PokemonsInLocation{}
	err = json.Unmarshal(data, &pokemonsInArea)
	if err != nil {
		return PokemonsInLocation{}, err
	}

	c.cache.Add(url, data)

	return pokemonsInArea, nil
}
