package commands

import (
	"encoding/json"
	"fmt"
)

var firstPageUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

func CommandMap(ctx *CliCommandCtx) error {
	if ctx.Next == nil && ctx.Previous == nil {
		ctx.Next = &firstPageUrl
	}
	url := *ctx.Next

	data, err := ctx.Cache.GetUrl(url)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	var areas LocationAreas
	if err := json.Unmarshal(data, &areas); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, row := range areas.Results {
		fmt.Println(row.Name)
	}

	ctx.Previous = areas.Previous
	ctx.Next = areas.Next

	return nil
}
