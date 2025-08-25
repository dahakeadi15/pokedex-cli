package pokeapi

import (
	"net/http"
	"time"

	"github.com/dahakeadi15/pokedex-cli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheDuration time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheDuration),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
