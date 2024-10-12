package commands

import (
	"fmt"
)

func CommandInspect(ctx *CliCommandCtx, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("missing pokemon name")
	}
	name := params[0]

	pokemon, ok := ctx.Pokemons[name]
	if !ok {
		return fmt.Errorf("you have not yet caught %s", name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
