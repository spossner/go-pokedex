package commands

import (
	"fmt"
)

func CommandPokedex(ctx *CliCommandCtx, params ...string) error {
	if len(ctx.Pokemons) == 0 {
		return fmt.Errorf("no pokemons yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range ctx.Pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
