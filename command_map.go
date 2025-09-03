package main

import (
	"errors"
	"fmt"

	"github.com/Predator792002/PokiInfo/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	url := pokeapi.BaseURL + "/location-area" // Default URL for first page
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	data, ok := cfg.pokeapiClient.Cache.Get(url)

	var locationsResp pokeapi.RespShallowLocations
	var err error
	if ok {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, data)
		if err != nil {
			return err
		}
	} else {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, nil)
		if err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	data, ok := cfg.pokeapiClient.Cache.Get(*cfg.prevLocationsURL)

	var locationsResp pokeapi.RespShallowLocations
	var err error
	if ok {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, data)
		if err != nil {
			return err
		}
	} else {
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, nil)
		if err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
