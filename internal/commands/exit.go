package commands

import "os"

func CommandExit(ctx *CliCommandCtx) error {
	os.Exit(0)
	return nil
}
