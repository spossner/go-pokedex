package commands

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

func CommandCatch(ctx *CliCommandCtx, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("missing pokemon name")
	}
	name := params[0]

	if _, ok := ctx.Pokemons[name]; ok {
		return fmt.Errorf("%s already caught", name)
	}

	fmt.Printf("trying to catch %s...\n", name)
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	data, err := ctx.Cache.GetUrl(url)
	if err != nil {
		return fmt.Errorf("error getting pokemon details: %w", err)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	exp := rand.IntN(220)
	fmt.Printf("Throwing a Pokeball (%d) at %s...\n", exp, name)
	if exp >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", name)
		ctx.Pokemons[name] = pokemon
	} else {
		fmt.Printf("%s (%d) escaped!\n", name, pokemon.BaseExperience)
	}

	return nil
}
