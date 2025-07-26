package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to your very own Pokedex!")
	fmt.Println("If you need help with commands type help.")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userCmd := cleanInput(scanner.Text())
		isCmd := isCommand(userCmd[0])
		if isCmd {
			commands := getCommands()
			err := commands[userCmd[0]].callback(cfg, userCmd)
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
	commands := getCommands()
	for _, command := range commands {
		if cmd == command.name {
			isCmd = true
			break
		}
	}
	return isCmd
}
