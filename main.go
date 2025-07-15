package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

// the function will take the string entered and break it up into seperate words
func cleanInput(text string) []string {
	return strings.Split(text, " ")
}
