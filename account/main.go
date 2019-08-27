package main

import (
	"fgame/fgame/account/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.Start(version, commit, date)
}
