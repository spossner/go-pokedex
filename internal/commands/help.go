package commands

import "fmt"

func CommandHelp(ctx *CliCommandCtx) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
