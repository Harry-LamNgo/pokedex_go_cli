# pokedex_go_cli

A command-line Pokédex built in Go. Explore Pokémon locations, discover creatures, and build your own virtual Pokédex — all from your terminal.

Data is fetched live from [PokéAPI](https://pokeapi.co/) and cached in memory for instant lookups while the program is running.

## Features

- **Interactive REPL** — type commands at a `Pokedex >` prompt
- **Location browsing** — navigate through Pokémon location areas
- **Explore** — see which Pokémon live in a given area
- **Catch** — throw a Pokéball at a Pokémon (it might escape!)
- **Inspect** — view height, weight, stats, and types of a caught Pokémon
- **Pokedex** — list every Pokémon you've caught this session
- **In-memory TTL cache** — API responses are cached so revisiting pages is instant

## Requirements

- Go 1.21 or newer
- An internet connection (to reach PokéAPI)

## Getting Started

```bash
git clone https://github.com/Harry-LamNgo/pokedex_go_cli.git
cd pokedex_go_cli
go run .
```

## Notes

- Use command `help` to show all available commands
- Caught Pokémon are stored in memory only and are lost when the program exits.
- All data comes from PokéAPI; no third-party Go dependencies are used.
