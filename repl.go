package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KrysPow/go_pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	arg                  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := strings.Fields(scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Printf(("Reading error, try a valid command, use help for commands."))
		}

		if _, ok := getCommand()[command[0]]; !ok {
			fmt.Printf("%s is not a command!\n", command[0])
			continue
		}

		if len(command) > 1 {
			conf.arg = &command[1]
		} else {
			conf.arg = nil
		}

		err := getCommand()[command[0]].callback(conf)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "Explores the given <area>, displays all pokemon occuring in that <area>. <area> can be an ID or the actual name",
			callback:    commandExplore,
		},
	}
}
