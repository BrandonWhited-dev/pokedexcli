package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fmt.Printf("Your command was: %s \n", cleanInput(scanner.Text())[0])
	}
}

// the function will take the string entered and break it up into seperate words
func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(text), " ")
}
