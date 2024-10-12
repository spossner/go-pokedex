package commands

import (
	"encoding/json"
	"fmt"
)

func CommandMapBack(ctx *CliCommandCtx) error {
	if ctx.Previous == nil {
		return fmt.Errorf("no such page")
	}
	url := *ctx.Previous

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
