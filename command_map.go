package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	cfg.nextLocationAreasURL = locationAreas.Next
	cfg.prevLocationAreasURL = locationAreas.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	cfg.nextLocationAreasURL = locationAreas.Next
	cfg.prevLocationAreasURL = locationAreas.Previous

	return nil
}
