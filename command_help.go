package main

import "fmt"

func commandHelp(*config, []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usuage:")
	fmt.Println()
	fmt.Println("anything in [] means a variable you enter")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Displays next location information")
	fmt.Println("mapb: Displays previous location information")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("explore [location name]: Displays the pokemon in an area")
	return nil
}
