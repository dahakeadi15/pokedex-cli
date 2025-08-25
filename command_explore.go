package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)
	locationArea, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid character") {
			return errors.New(" - No pokemon found in the area")
		}
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
