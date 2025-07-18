package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userCmd := cleanInput(scanner.Text())[0]
		isCmd := isCommand(userCmd)
		if isCmd {
			err := commands[userCmd].callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// the function will take the string entered and break it up into seperate words
func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(text), " ")
}

// checks if the string is a command
func isCommand(cmd string) bool {
	isCmd := false
	for _, command := range commands {
		if cmd == command.name {
			isCmd = true
			break
		}
	}
	return isCmd
}
