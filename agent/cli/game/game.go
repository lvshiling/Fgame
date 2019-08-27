package game

import (
	"github.com/urfave/cli"
)

var (
	cmds = []cli.Command{
		configCommand,
	}
)

var (
	GameCommand = cli.Command{
		Name:        "game",
		Usage:       "agent [global options] game [command options] [arguments...].",
		Description: "game agent",
		Subcommands: cmds,
	}
)
