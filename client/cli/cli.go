package cli

import (
	"encoding/json"

	"fgame/fgame/client/game"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"
)

var (
	debug      = false
	deviceMac  = ""
	configFile = ""
)

func Start() {
	app := cli.NewApp()
	app.Name = "game client"
	app.Usage = "client [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "device",
			Value:       deviceMac,
			Usage:       "device",
			Destination: &deviceMac,
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

func initConfig(configFile string) (sc *game.ClientOptions, err error) {
	log.Infoln("正在读取配置")
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &game.ClientOptions{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}
	log.Infoln("读取配置成功")
	return sc, nil
}

func before(c *cli.Context) error {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./logs/info.log",
		log.InfoLevel:  "./logs/info.log",
		log.WarnLevel:  "./logs/info.log",
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

func start(ctx *cli.Context) {
	log.Infoln("正在启动客户端")

	sc, err := initConfig(configFile)
	if err != nil {
		log.Fatalln("init config file failed ", err)
	}

	err = game.InitGame(sc)
	if err != nil {
		log.Fatalln("init client failed ", err)
	}

}

//TODO 关闭
