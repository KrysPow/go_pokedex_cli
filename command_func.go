package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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

func commandExplore(conf *config) error {
	if conf.arg == nil {
		return fmt.Errorf("No location area given to explore")
	}

	locAreaDetails, err := conf.pokeapiClient.ListPokemonInArea(conf.arg)
	if err != nil {
		return err
	}

	for _, encounter := range locAreaDetails.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *config) error {
	pokemonDetails, err := conf.pokeapiClient.PokemonDetails(conf.arg)
	if err != nil {
		return err
	}

	fmt.Printf("Throw a Pokeball at %s ", pokemonDetails.Name)
	for i := 1; i < 4; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * time.Duration(i) * 30)
	}
	fmt.Println("")

	chances := calculateCatchingChance(pokemonDetails.BaseExperience)
	if chances < 0.5 {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
		conf.caughtPokemon[pokemonDetails.Name] = pokemonDetails
	}
	return nil
}

func calculateCatchingChance(base_exp int) float64 {
	random := rand.Intn(base_exp)
	return float64(random) / float64(base_exp)
}
