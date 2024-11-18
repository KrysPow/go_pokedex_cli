package main

import (
	"time"

	"github.com/KrysPow/go_pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Hour, 5*time.Second)
	conf := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(conf)
}
