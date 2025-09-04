package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(pokeName string) ([]string, error) {
	url := BaseURL + "/location-area" + "/" + pokeName

	var dat []byte
	cacheData, ok := c.Cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 200:
			dat, err = io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			c.Cache.Add(url, dat)
		case 404:
			return nil, fmt.Errorf("pokemon Not found")
		default:
			return nil, fmt.Errorf("error producing response")
		}

	} else {
		dat = cacheData
	}

	pokemonEncountersResp := Encounters{}
	err := json.Unmarshal(dat, &pokemonEncountersResp)
	if err != nil {
		return nil, err
	}

	var pokemons []string

	for _, enc := range pokemonEncountersResp.PokemonEncounters {
		pokemons = append(pokemons, enc.Pokemon.Name)
	}

	return pokemons, nil

}
