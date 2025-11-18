package pokeapi

import (
	"net/http"
	"time"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/cache"
)

// Client -
type Client struct {
	httpClient   http.Client
	pokeapiCache *cache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeapiCache: cache.NewCache(timeout),
	}
}
