package pokeapi

import (
	"net/http"
	"time"

	"github.com/Predator792002/PokiInfo/internal/pokecache"
)

type Client struct {
	Cache      *pokecache.Cache // Add this field
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client { // Two parameters now
	return Client{
		Cache: pokecache.NewCache(cacheInterval), // Create cache here
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
