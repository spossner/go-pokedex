package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spossner/pokedexcli/internal/commands"
	"github.com/spossner/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ctx := &commands.CliCommandCtx{
		Cache: pokecache.NewCache(3 * time.Minute),
	}

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
