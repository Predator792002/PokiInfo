package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cacheData []byte) (RespShallowLocations, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var dat []byte

	if cacheData == nil {
		// Make HTTP request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		// Add to cache after successful request
		c.Cache.Add(url, dat)
	} else {
		// Use cached data
		dat = cacheData
	}

	// Unmarshal once (whether from cache or HTTP)
	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
