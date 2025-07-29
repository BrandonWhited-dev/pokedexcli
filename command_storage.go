package main

import (
	"errors"
	"fmt"
)

func commandStorage(cfg *config, args []string) error {
	if len(cfg.Storage) < 1 {
		return errors.New("You have not caught any pokemon yet.")
	}

	fmt.Println("Your Pokemon:")
	for name, _ := range cfg.Storage {
		fmt.Printf("   - %s \n", name)
	}

	return nil
}
