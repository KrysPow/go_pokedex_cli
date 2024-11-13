package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exits the Pokedex\n")
	return nil
}

func commandExit() error {
	return fmt.Errorf("Exit")
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	{
		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)

		exitLoop := false
		for exitLoop == false {
			fmt.Print("Pokedex > ")
			for scanner.Scan() {
				input := scanner.Text()
				if _, ok := commands[input]; !ok {
					fmt.Printf("%s is not a command!\n", input)
					break
				}

				err := commands[input].callback()
				if err != nil {
					exitLoop = true
				}
				break
			}
		}
	}

}
