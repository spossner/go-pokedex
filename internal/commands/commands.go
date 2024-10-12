package commands

import "github.com/spossner/pokedexcli/internal/pokecache"

type CliCommandCtx struct {
	Next, Previous *string
	Cache          *pokecache.Cache
}

type CliCommand struct {
	Name, Description string
	Fn                func(ctx *CliCommandCtx) error
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
	}
}

func GetCommand(cmd string) (CliCommand, bool) {
	CliCommand, ok := GetCommands()[cmd]
	return CliCommand, ok
}
