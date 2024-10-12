package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var firstPageUrl string = "https://pokeapi.co/api/v2/location-area"

func CommandMap(ctx *CliCommandCtx) error {
	if ctx.Next == nil && ctx.Previous == nil {
		ctx.Next = &firstPageUrl
	}
	res, err := http.Get(*ctx.Next)
	if err != nil {
		return fmt.Errorf("error getting next page: %w", err)
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
