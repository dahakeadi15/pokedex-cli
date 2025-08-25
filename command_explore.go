package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)
	pokemonInArea, err := cfg.pokeapiClient.ListPokemonInArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonInArea {
		fmt.Println(" - ", pokemon)
	}

	return nil
}
