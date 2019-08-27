package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"

	"github.com/codegangsta/negroni"
	"github.com/xozrc/pkg/osutils"
)

var (
	debug      = false
	configFile = ""
)

func Start() {

	app := cli.NewApp()
	app.Name = "resource_server"
	app.Usage = "resource_server [global options]"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug ",
			Destination: &debug,
			EnvVar:      "DEBUG",
		},
		cli.StringFlag{
			Name:        "config,c",
			Value:       configFile,
			Usage:       "config file",
			Destination: &configFile,
		},
	}

	app.Before = before
	app.Action = start
	app.Run(os.Args)
}

func before(c *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.ErrorLevel: "./logs/error.log",
	}))

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	return nil
}

type serverConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func newServerConfig(configFile string) (sc *serverConfig, err error) {
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	sc = &serverConfig{}
	err = json.Unmarshal(c, sc)
	if err != nil {
		return nil, err
	}
	return
}

func start(ctx *cli.Context) {
	var sc *serverConfig

	config, err := filepath.Abs(configFile)
	if err != nil {
		log.Fatalln("filepath abs failed:", err)
	}
	tsc, err := newServerConfig(config)
	if err != nil {
		log.Fatalln("read config file failed:", err)
	}
	sc = tsc

	n := negroni.Classic()

	addr := fmt.Sprintf("%s:%d", sc.Host, sc.Port)

	//register interruput handler
	hooker := osutils.NewInterruptHooker()
	hooker.AddHandler(osutils.InterruptHandlerFunc(stop))
	log.Println("trying to listen ", addr)
	n.Run(addr)
}

func stop() {
	log.Println("stop server")
}
