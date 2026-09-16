package main

import (
	"time"

	"github.com/Harry-LamNgo/pokdex_go_cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewCilent(5 * time.Second)
	// Initialize Config (cfg) -- Get all supported command of Pokedex
	cfg := config{
		commands:      getSupportedCommands(),
		pokeapiClient: pokeClient,
	}
	runREPL(&cfg)
}
