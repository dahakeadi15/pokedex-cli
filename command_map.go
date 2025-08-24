package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
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
