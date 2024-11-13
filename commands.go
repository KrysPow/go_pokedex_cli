package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exits the Pokedex\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
