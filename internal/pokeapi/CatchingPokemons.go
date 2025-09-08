package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemons(name string) (CaughtPokemon, error) {
	url := BaseURL + "/pokemon/" + name
	if b, ok := c.Cache.Get(url); ok {
		cp := CaughtPokemon{}
		if err := json.Unmarshal(b, &cp); err != nil {
			return CaughtPokemon{}, err
		}
		return cp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CaughtPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CaughtPokemon{}, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return CaughtPokemon{}, err
		}
		c.Cache.Add(url, b)
		var cp CaughtPokemon
		if err := json.Unmarshal(b, &cp); err != nil {
			return CaughtPokemon{}, err
		}
		return cp, nil
	case 404:
		return CaughtPokemon{}, fmt.Errorf("pokemon not found")
	default:
		return CaughtPokemon{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
}
