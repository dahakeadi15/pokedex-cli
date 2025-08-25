package main

import (
	"time"

	"github.com/dahakeadi15/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	pokemon              map[string]pokeapi.Pokemon
}

func main() {
	apiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: apiClient,
		pokemon:       map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
