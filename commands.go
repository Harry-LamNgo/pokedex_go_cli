package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
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

		"pokedex": {
			name:        "pokedex",
			description: "Displays list of all captured Pokemon",
			callback:    commandPokedex,
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

		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Displays captured pokemon information (stats, types, ...)",
			callback:    commandInspect,
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

	// Tracking location and Pokemon encountered in that location
	cfg.lastArea = listPokemonResp.Name
	cfg.encountered = make(map[string]bool, len(listPokemonResp.PokemonEncounters))

	fmt.Printf("Exploring %s ...\n", listPokemonResp.Name)
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounters := range listPokemonResp.PokemonEncounters {
		cfg.encountered[pokemonEncounters.Pokemon.Name] = true
		fmt.Println(" - " + pokemonEncounters.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("catch command requires one name of pokemon to execute")
	}

	targetPokemon := cfg.args[0]

	if cfg.lastArea == "" {
		return errors.New("explore an area first with the 'explore' command")
	}

	if !cfg.encountered[targetPokemon] {
		return fmt.Errorf("%s is not found in %s", targetPokemon, cfg.lastArea)
	}

	pokemonResp, err := cfg.pokeapiClient.FetchTargetPokemon(targetPokemon)
	if err != nil {
		return err
	}

	// Chance to catch Pokemon --> convert to 1.0 float scale

	// chance stays in [~0.14, ~0.7] for every real baseExp value
	// Ex: Blissey has baseExp = 635 -> catch chance = 1 / (1 + 635/100) = ~0.13

	catchChance := func(baseExp int) float64 {
		return 1.0 / (1.0 + float64(baseExp)/100.0)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", targetPokemon)

	if rand.Float64() < catchChance(pokemonResp.BaseExperience) {
		cfg.caughtpokemon[targetPokemon] = pokemonResp
		fmt.Printf("%v was caught!\n", pokemonResp.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped!\n", pokemonResp.Name)
	}

	return nil
}

func commandInspect(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("inspect command requires one name of captured pokemon to execute")
	}

	inspectedPokemon := cfg.caughtpokemon[cfg.args[0]]

	if inspectedPokemon.Name == "" {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf(`
Name: %v
Height: %d
Weight: %d
`, inspectedPokemon.Name, inspectedPokemon.Height, inspectedPokemon.Weight)

	fmt.Println("Stats:")
	for _, stats := range inspectedPokemon.Stats {
		fmt.Printf(" - %v: %v\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketypes := range inspectedPokemon.Types {
		fmt.Printf(" - %v\n", poketypes.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config) error {
	if len(cfg.caughtpokemon) == 0 {
		fmt.Println("Your Pokedex is empty - Let explore and catch some pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.caughtpokemon {
		fmt.Printf(" - %v\n", pokemon.Name)
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
