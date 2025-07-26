package main

import "fmt"

func commandHelp(*config, []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usuage:")
	fmt.Println()
	fmt.Println("anything in [] means a variable you enter")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
