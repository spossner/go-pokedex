package commands

import (
	"time"

	"github.com/spossner/pokedexcli/internal/pokecache"
)

type CliCommandCtx struct {
	Next, Previous *string
	Pokemons       map[string]Pokemon
	Cache          *pokecache.Cache
}

func NewCliCommandCtx() *CliCommandCtx {
	return &CliCommandCtx{
		Pokemons: make(map[string]Pokemon),
		Cache:    pokecache.NewCache(3 * time.Minute),
	}
}

type CliCommand struct {
	Name, Description string
	Fn                func(ctx *CliCommandCtx, params ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Fn:          CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Fn:          CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of the next 20 location areas in the Pokemon world",
			Fn:          CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of previous 20 location areas in the Pokemon world",
			Fn:          CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the pokemon in a given area",
			Fn:          CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Tries to catch a given pokemon",
			Fn:          CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows the stats of a caught pokemon",
			Fn:          CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays a list of all pokemons already caugth",
			Fn:          CommandPokedex,
		},
	}
}

func GetCommand(cmd string) (CliCommand, bool) {
	CliCommand, ok := GetCommands()[cmd]
	return CliCommand, ok
}
