package commands

import "os"

func CommandExit(ctx *CliCommandCtx, params ...string) error {
	os.Exit(0)
	return nil
}
