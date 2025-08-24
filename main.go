package main

import (
	"time"

	"github.com/dahakeadi15/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
}

func main() {
	apiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: apiClient,
	}
	startRepl(cfg)
}
