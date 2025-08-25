package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid character") {
			return errors.New("no such pokemon exist in the database")
		}
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonBaseXp := pokemonInfo.BaseExperience
	chance := rand.Intn(pokemonBaseXp)

	percentChance := (float64(chance) / float64(pokemonBaseXp)) * 100.0

	if percentChance > 60.0 {
		cfg.pokemon[pokemonName] = pokemonInfo
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
