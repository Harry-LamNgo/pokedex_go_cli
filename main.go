package main

func main() {
	// Initialize Config (cfg) -- Get all supported command of Pokedex
	cfg := config{
		commands: getSupportedCommands(),
	}
	runREPL(&cfg)
}
