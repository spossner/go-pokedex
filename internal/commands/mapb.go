package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CommandMapBack(ctx *CliCommandCtx) error {
	if ctx.Previous == nil {
		return fmt.Errorf("no such page")
	}
	res, err := http.Get(*ctx.Previous)
	if err != nil {
		return fmt.Errorf("error getting previous page: %w", err)
	}
	defer res.Body.Close()

	var areas LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&areas); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, row := range areas.Results {
		fmt.Println(row.Name)
	}

	ctx.Previous = areas.Previous
	ctx.Next = areas.Next

	return nil
}
