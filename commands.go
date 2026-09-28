package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/Harry-LamNgo/pokdex_go_cli/internal/pokeapi"
)

// The structure of command - it must have "name", "description" and "callback" function
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Method wrap commands in a function --> this function only run after main() run
//   --> re-create map of commands each time this function call -> cannot accidentally delete the command like declare only one time by global var

func getSupportedCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"explore": {
			name:        "explore <target-location-area>",
			description: "Displays a list of all Pokémon encountered in the target location area.",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch <pokemon-name>",
			description: "Atempt to catch a pokemon",
			callback:    commandCatch,
		},

		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// os.Exit with code zero (0) -- indicates success
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.commands {
		fmt.Printf("	%s - %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	// Do the Fetch Locations from pokeAPI -> with nextURL (if nextURL is nil -> FetchLocations handles that using base URL)
	locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	// Track the Previous URL and Next URL
	cfg.previousURL = locationResp.Previous
	cfg.nextURL = locationResp.Next

	// Iterate and print location-area
	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	// Check the previous URL exist
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.previousURL)
	if err != nil {
		return err
	}

	// Track the Previous URL and Next URL
	cfg.previousURL = locationResp.Previous
	cfg.nextURL = locationResp.Next

	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("explore command requires one specific location-area name to execute")
	}

	areaName := cfg.args[0]
	// fmt.Printf("Debug areaName - %q\n", areaName)
	listPokemonResp, err := cfg.pokeapiClient.FetchPokemons(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s ...\n", listPokemonResp.Name)
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounters := range listPokemonResp.PokemonEncounters {
		fmt.Println(" - " + pokemonEncounters.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("catch command requires one name of pokemon to execute")
	}

	targetPokemon := cfg.args[0]

	targetPokemonResp, err := cfg.pokeapiClient.FetchTargetPokemon(targetPokemon)
	if err != nil {
		return err
	}

	pokemonName := targetPokemonResp.Name

	// Chance to catch Pokemon --> convert to 1.0 float scale

	// chance stays in [~0.1, ~0.9] for every real baseExp value
	// Ex: Blissey has baseExp = 635 -> catch chance = 1 / (1 + 635/100) = ~0.13

	catchChance := func(baseExp int) float64 {
		return 1.0 / (1.0 + float64(baseExp)/100.0)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", targetPokemon)

	if rand.Float64() < catchChance(targetPokemonResp.BaseExperience) {
		cfg.caughtpokemon[pokemonName] = pokeapi.CaughtPokemon{
			Name: pokemonName,
		}

		fmt.Printf("%v was caught!\n", pokemonName)
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}

// Method declare Global var -- Go auto runs init() before main() -> add all commands into one map and shared everywhere in package main
//   --> this map lives for entire duration of the program
/*
var supportedCommands = map[string]cliCommand{}

func init() {
	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}
*/
