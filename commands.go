package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, _ := range getCommand() {
		fmt.Printf("%s: %s\n", getCommand()[k].name, getCommand()[k].description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
