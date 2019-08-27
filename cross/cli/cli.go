package cli

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"
)

var (
	debug      = false
	configFile = ""
	commands   = []cli.Command{}
)

func Start() {
	app := cli.NewApp()
	app.Name = "cross"
	app.Usage = "cross [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
	}
	app.Before = before
	app.Action = crossServerStart
	app.Commands = commands
	app.Run(os.Args)
}

func before(c *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.InfoLevel:  "./logs/info.log",
		log.WarnLevel:  "./logs/warn.log",
		log.ErrorLevel: "./logs/error.log",
		log.FatalLevel: "./logs/error.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}
