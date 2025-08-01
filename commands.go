package main

// command struct
type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

// map of commands
func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
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
		"explore": {
			name:        "explore",
			description: "Type a location name after explore. Displays pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Enter the pokemons name after the command. Allows you to attempt to catch a pokemon.",
			callback:    commandCatch,
		},
		"storage": {
			name:        "storage",
			description: "Displays all pokemon you have caught.",
			callback:    commandStorage,
		},
		"delete": {
			name:        "delete",
			description: "Deletes save data",
			callback:    commandDelete,
		},
		"inspect": {
			name:        "inspect",
			description: "Enter a pokemons name after the command. Displays data about pokemon.",
			callback:    commandInspect,
		},
	}
	return commands
}
