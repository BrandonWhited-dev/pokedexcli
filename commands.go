package main

// command struct
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// map of commands
var commands map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays Next locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays Previous locations",
		callback:    commandMapb,
	},
}
