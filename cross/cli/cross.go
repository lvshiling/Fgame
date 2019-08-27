package cli

import (
	"encoding/json"
	"fgame/fgame/cross/server"
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func readCrossServerConfig(configFile string) (sc *server.CrossGameServerOptions, err error) {
	log.Infoln("正在读取配置")
	abs, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}
	sc = &server.CrossGameServerOptions{}
	if err = json.Unmarshal(bs, sc); err != nil {
		return nil, err
	}
	log.Infoln("读取配置成功")
	return sc, nil
}

func crossServerStart(c *cli.Context) {
	log.Infoln("正在启动跨服服务器")

	sc, err := readCrossServerConfig(configFile)
	if err != nil {
		log.Fatalln("读取跨服服务器配置失败", err)
	}

	gs, err := server.InitCrossGameServer(sc)
	if err != nil {
		log.Fatalln("初始化跨服服务器失败", err)
	}
	err = gs.Start()
	if err != nil {
		log.Fatalln("启动跨服服务器", err)
	}
}
