package main

import "fmt"

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usuage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Displays next location information")
	fmt.Println("map: Displays previous location information")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
