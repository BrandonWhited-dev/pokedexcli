package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	WriteData(cfg)
	os.Exit(0)
	return nil
}
