package cli

import (
	"encoding/json"
	"fgame/fgame/logserver/server"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/urfave/cli"
)

var (
	debug      = false
	configFile = ""
)

func Start() {
	app := cli.NewApp()
	app.Name = "log server"
	app.Usage = "log [global options] command [command options] [arguments...]."

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
	app.Action = start
	app.Run(os.Args)
}

func initConfig(configFile string) (sc *server.LogServerOptions, err error) {
	log.Infoln("正在读取配置")
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &server.LogServerOptions{}
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

func start(c *cli.Context) {
	log.Infoln("正在启动日志服务器")

	sc, err := initConfig(configFile)
	if err != nil {
		log.Fatalln("init config file failed ", err)
	}

	err = server.Init(sc)
	if err != nil {
		log.Fatalln("启动日志服务器失败", err)
	}

}
