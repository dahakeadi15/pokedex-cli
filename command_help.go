package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	commands := getCommands()
	longestLength := 0
	for k := range commands {
		longestLength = max(len(k), longestLength)
	}

	for _, cmd := range commands {
		name := cmd.name + ":"
		for i := len(name); i < longestLength+1; i++ {
			name += " "
		}
		fmt.Printf("%s %s\n", name, cmd.description)
	}
	fmt.Println("")

	return nil
}
