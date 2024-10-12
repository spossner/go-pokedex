package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spossner/pokedexcli/internal/commands"
)

func main() {
	ctx := commands.NewCliCommandCtx()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("an error occured")
			commands.CommandExit(ctx)
		}

		ops := strings.Split(scanner.Text(), " ")

		if cmd, ok := commands.GetCommand(ops[0]); ok {
			if err := cmd.Fn(ctx, ops[1:]...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknwon Command", scanner.Text())
		}
	}
}
