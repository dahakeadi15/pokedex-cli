package main

import "fmt"

func commandExplore(cfg *config, locationArea string) error {
	fmt.Printf("Exploring %s...\n", locationArea)

	pokemonInArea, err := cfg.pokeapiClient.ListPokemonInArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonInArea {
		fmt.Println(" - ", pokemon)
	}

	return nil
}
