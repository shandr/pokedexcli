package pokeapi

import (
	"github.com/shandr/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
