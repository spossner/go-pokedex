package commands

import "fmt"

func CommandHelp(ctx *CliCommandCtx, params ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
