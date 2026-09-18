package main

import (
	"time"

	"github.com/Harry-LamNgo/pokdex_go_cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Minute, 5*time.Second)
	// Initialize Config (cfg) -- Get all supported command of Pokedex
	cfg := config{
		commands:      getSupportedCommands(),
		pokeapiClient: pokeClient,
	}
	runREPL(&cfg)
}
