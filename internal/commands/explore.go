package commands

import (
	"encoding/json"
	"fmt"
)

func CommandExplore(ctx *CliCommandCtx, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("missing area name")
	}
	location := params[0]
	fmt.Printf("Exploring %s...\n", location)
	url := "https://pokeapi.co/api/v2/location-area/" + location
	data, err := ctx.Cache.GetUrl(url)
	if err != nil {
		return fmt.Errorf("error getting location area details: %w", err)
	}

	var area LocationArea
	if err := json.Unmarshal(data, &area); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	if len(area.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon found")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range area.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}
