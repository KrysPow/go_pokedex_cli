package main

import (
	"fmt"
	"os"
)

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println("")
	for k := range getCommand() {
		fmt.Printf("%s: %s\n",
			getCommand()[k].name,
			getCommand()[k].description)
	}
	return nil
}

func commandExit(conf *config) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *config) error {
	listLoc, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = listLoc.Next
	conf.previousLocationsURL = listLoc.Previous

	for _, result := range listLoc.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(conf *config) error {
	if conf.previousLocationsURL == nil {
		return fmt.Errorf("you are on the first page")
	}

	listLoc, err := conf.pokeapiClient.ListLocations(conf.previousLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = listLoc.Next
	conf.previousLocationsURL = listLoc.Previous

	for _, result := range listLoc.Results {
		fmt.Println(result.Name)
	}
	return nil
}
