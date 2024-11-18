package pokeapi

import (
	"net/http"
	"time"

	"github.com/KrysPow/go_pokedex/internal/pokecache"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient
func NewClient(internal time.Duration, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(internal),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
