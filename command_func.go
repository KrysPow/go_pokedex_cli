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

func commandInspect(conf *config) error {
	if conf.arg == nil {
		return fmt.Errorf("A pokemon must be given in combination with <inspect>")
	}
	if pokemonDetails, ok := conf.caughtPokemon[*conf.arg]; !ok {
		return fmt.Errorf("A pokemon that was not caught cannot be inspected. If you used a pokemons ID, use its name to inspect it")
	} else {
		fmt.Printf("Name: %s\n", pokemonDetails.Name)
		fmt.Printf("Height: %v\n", pokemonDetails.Height)
		fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemonDetails.Stats {
			fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types")
		for _, typ := range pokemonDetails.Types {
			fmt.Printf(" - %v\n", typ.Type.Name)
		}
	}
	return nil
}

func commandPokedex(conf *config) error {
	if len(conf.caughtPokemon) == 0 {
		return fmt.Errorf("No pokemon caught")
	}
	for pokemon := range conf.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
