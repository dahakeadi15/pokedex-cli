package main

import (
	"time"

	"github.com/dahakeadi15/pokedex-cli/internal/pokeapi"
	"github.com/dahakeadi15/pokedex-cli/internal/pokecache"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	pokecache            pokecache.Cache
}

func main() {
	apiClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: apiClient,
		pokecache:     cache,
	}
	startRepl(cfg)
}
