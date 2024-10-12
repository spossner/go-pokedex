package main

import (
	"bufio"
	"fmt"
	"os"
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

		if cmd, ok := commands.GetCommand(scanner.Text()); ok {
			if err := cmd.Fn(ctx); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknwon Command", scanner.Text())
		}
	}
}
