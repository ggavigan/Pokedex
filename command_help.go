package main

import "fmt"

func commandHelp(*config) error {
	fmt.Println()
	fmt.Println("Pokedex commands help:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
